package tailf

import (
	"fmt"
	"go_dev/day11/logagent/conf"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

const (
	StatusNormal = 1
	StatusDelete = 2
)

type TailObj struct {
	tail        *tail.Tail     // each tail reads  a log file
	collectConf conf.Collector // config in ini file
	exitChan    chan int       // exit channel
	status      int
}

// message send to kafka, inclues msg, and topic
type TextMsg struct {
	Msg   string
	Topic string
}

// we have several log files, each log file has a *tail*
type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan  chan *TextMsg
	lock     sync.Mutex
}

var (
	tailObjMgr *TailObjMgr
)

func GetOneLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return
}

func InitTailF(collectors []conf.Collector) (err error) {

	if len(collectors) == 0 {
		logs.Warn("no collector definied!")
	}

	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, 100),
	}

	for _, v := range collectors {
		createNewTask(v)
	}
	return
}

func readFromTail(tailObj *TailObj) {
	for true {
		select {
		case line, ok := <-tailObj.tail.Lines:
			if !ok {
				logs.Warn("tail file close reopen, filename:%s\n", tailObj.tail.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			textMsg := &TextMsg{
				Msg:   line.Text,
				Topic: tailObj.collectConf.Topic,
			}
			tailObjMgr.msgChan <- textMsg
		case <-tailObj.exitChan:
			logs.Warn("tail object exits, conf: %v", tailObj.collectConf)
			return
		}
	}
}

// update collector config, when receives a config update
func UpdateConfig(confs []conf.Collector) {
	// add lock
	tailObjMgr.lock.Lock()
	defer tailObjMgr.lock.Unlock()

	for _, oneConf := range confs {
		var isRunning = false
		for _, v := range tailObjMgr.tailObjs {
			if oneConf.LogPath == v.collectConf.LogPath {
				isRunning = true
				break
			}
		}
		if isRunning {
			continue
		}
		logs.Debug("create a new tail task, %v", oneConf)
		createNewTask(oneConf)
	}

	// delete and save exists tasks!
	var tailObjs []*TailObj
	for _, obj := range tailObjMgr.tailObjs {
		obj.status = StatusDelete
		for _, oneConf := range confs {
			if oneConf.LogPath == obj.collectConf.LogPath {
				obj.status = StatusNormal
				break
			}
		}
		if obj.status == StatusDelete {
			obj.exitChan <- 1 // delete task
			continue
		}
		tailObjs = append(tailObjs, obj)
	}
	tailObjMgr.tailObjs = tailObjs // save existing tailObjs
	return
}

func createNewTask(conf conf.Collector) {
	tails, tailErr := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen: true,
		Follow: true,
		// Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if tailErr != nil {
		fmt.Println("tail file err:", tailErr)
		logs.Error("collect filename[%s] failed, err:%v", conf.LogPath, tailErr)
		return
	}

	tailObj := &TailObj{
		tail:        tails,
		collectConf: conf,
		exitChan:    make(chan int, 1),
	}

	tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, tailObj)
	go readFromTail(tailObj)
}
