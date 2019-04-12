package main

import "fmt"

type userinfo map[string]map[string]string

func testMap() {
	// var m map[string]string
	// m = make(map[string]string, 10) // 申请内存空间

	m := make(map[string]string, 10) // 申请内存空间

	m["aa"] = "AA"
	m["bb"] = "AA"
	fmt.Println(m)

	var b map[string]string = map[string]string{}
	b["a"] = "A"
	b["b"] = "BBBBBBBBBBBBBB"
	fmt.Println(len(b))
}

func testMap2() {
	// map 嵌套两层
	a := make(map[string]map[string]string, 10)
	a["a"] = make(map[string]string) // 申请空间
	a["a"]["key1"] = "v1"
	a["a"]["key2"] = "v2"
	a["a"]["key3"] = "v3"
	fmt.Println(a)
}

func modify(m userinfo) {
	_, ok := m["zs"]
	if !ok {
		m["zs"] = make(map[string]string)
	}
	m["zs"]["pwd"] = "213"
	m["zs"]["nickname"] = "小龙虾"
}

func testMap3() {
	a := make(userinfo, 100)
	modify(a)
	fmt.Println(a)
}

// 遍历map
func trav(m map[string]string) {
	for k, v := range m {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
}

func testMap4() {
	m := make(map[string]string, 10) // 申请内存空间

	m["a"] = "AA"
	m["b"] = "AA"
	m["c"] = "AA"
	m["d"] = "AA"
	m["e"] = "AA"
	m["f"] = "AA"
	m["g"] = "AA"
	m["h"] = "AA"
	m["i"] = "AA"
	m["j"] = "AA"
	m["k"] = "AA"
	m["l"] = "AA"
	m["m"] = "AA"

	// 遍历
	trav(m)
	fmt.Println(len(m))
	delete(m, "h")
	delete(m, "z")
	fmt.Println(len(m))

}

func main() {
	// testMap()
	// testMap2()
	// testMap3()
	testMap4()
}
