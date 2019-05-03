package tailf

import (
	"fmt"
	"go_dev/day11/logagent/conf"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

type TailObj struct {
	tail        *tail.Tail     // each tail reads  a log file
	collectConf conf.Collector // config in ini file
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
		panic("no collector definied!")
	}

	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, 100),
	}

	for _, v := range collectors {
		tails, tailErr := tail.TailFile(v.LogPath, tail.Config{
			ReOpen:    true,
			Follow:    true,
			Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
			MustExist: false,
			Poll:      true,
		})
		if tailErr != nil {
			fmt.Println("tail file err:", tailErr)
			err = tailErr
			return
		}

		tailObj := &TailObj{
			tail:        tails,
			collectConf: v,
		}

		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, tailObj)
		go readFromTail(tailObj)
	}
	return
}

func readFromTail(tailObj *TailObj) {
	for true {
		line, ok := <-tailObj.tail.Lines
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
	}
}
