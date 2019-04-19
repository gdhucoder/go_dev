package balance

import (
	"fmt"
	"math/rand"
	"strconv"
)

// 一致性哈希算法 负载均衡
type HashBalancer struct {
}

// 注册 一致性hash负载均衡
func init() {
	RegisterBalancer("hash", &HashBalancer{})
}

// 实现DoBalance
func (p *HashBalancer) DoBalance(insts []*Instance) (inst *Instance, err error) {
	var defaultKey = fmt.Sprintf("%d", rand.Intn(5000))
	var num, _ = strconv.Atoi(defaultKey)
	var lens = len(insts)
	inst = insts[num%lens]
	return
}
