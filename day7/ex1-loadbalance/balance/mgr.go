package balance

import "fmt"

type BalanceMgr struct {
	allBalancer map[string]Balancer
}

var mgr = BalanceMgr{
	allBalancer: make(map[string]Balancer),
}

func (p *BalanceMgr) registerBalancer(name string, balancer Balancer) {
	p.allBalancer[name] = balancer
}

func RegisterBalancer(name string, balancer Balancer) {
	mgr.registerBalancer(name, balancer)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = fmt.Errorf("Not found %s balancer", name)
		return
	}
	inst, err = balancer.DoBalance(insts)
	fmt.Printf("using %s balance.\n", name)
	return
}
