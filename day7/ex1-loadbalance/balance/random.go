package balance

import (
	"errors"
	"math/rand"
)

func init() {
	RegisterBalancer("random", NewRandom())
}

type Random struct {
}

func (r *Random) DoBalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no instances")
		return
	}
	numOfInsts := len(insts)
	idx := rand.Intn(numOfInsts)
	inst = insts[idx]
	return
}

func NewRandom() *Random {
	return &Random{}
}
