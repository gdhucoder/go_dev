package main

import (
	"fmt"
	"sort"
)

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

func testMap5() {
	// 切片中的元素是map
	var a []map[int]int
	a = make([]map[int]int, 5) // 初始化
	if a[0] == nil {
		a[0] = make(map[int]int, 5)
	}
	a[0][1] = 2
	fmt.Println(a)
	// [map[1:2] map[] map[] map[] map[]]
}

func testMapSort() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[0] = 0
	a[1] = 0
	a[2] = 0
	a[3] = 0
	a[4] = 0

	var keys []int

	// map 中的key是无序的
	for k1, v := range a {
		fmt.Println(k1, v)
		keys = append(keys, k1)
	}

	sort.Ints(keys)

	for v := range keys {
		fmt.Println(v, a[v])
	}

	// 未排序
	// 1 0
	// 2 0
	// 3 0
	// 4 0
	// 0 0

	// 排序后
	// 0 0
	// 1 0
	// 2 0
	// 3 0
	// 4 0

}

func testMapReverse() {
	var a map[string]int
	a = make(map[string]int, 5)
	a["adf"] = 123
	a["ddd"] = 22

	var b map[int]string
	b = make(map[int]string, 5)
	for k, v := range a {
		b[v] = k
	}
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	// testMap()
	// testMap2()
	// testMap3()
	// testMap4()
	// testMap5()
	// testMapSort()
	testMapReverse()
}
