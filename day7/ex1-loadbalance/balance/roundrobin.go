package balance

import "errors"

func init() {
	RegisterBalancer("roundrobin", NewRoundRobin())
}

type RoundRobin struct {
	curInstIdx int
}

func (r *RoundRobin) DoBalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no instances")
		return
	}
	idx := r.curInstIdx % len(insts)
	inst = insts[idx]
	r.curInstIdx++
	return
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{
		curInstIdx: 0,
	}
}
