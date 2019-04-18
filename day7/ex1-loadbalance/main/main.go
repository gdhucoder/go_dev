package main

import (
	"fmt"
	"go_dev/day7/ex1-loadbalance/balance"
	"math/rand"
	"time"
)

func main() {
	var insts []*balance.Instance // 空的切片
	for i := 0; i < 2; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(256), rand.Intn(256))
		// port := 5000 + rand.Intn(100)
		port := 5000
		var inst *balance.Instance
		inst = balance.NewInstance(host, port)
		insts = append(insts, inst)
	}

	fmt.Println(insts)

	// var roundRobin *balance.RoundRobin
	// roundRobin = balance.NewRoundRobin()
	balancer := balance.NewRoundRobin() // 负载均衡
	for i := 0; i < 20; i++ {

		inst, err := balancer.DoBalance(insts)

		if err != nil {
			fmt.Println("balance error: ", err)
			return
		}

		fmt.Println(inst)
		time.Sleep(time.Second)
	}

}
