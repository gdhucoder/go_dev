package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"username"`
	NickName string
	Age      int
	Phone    string
}

func testMap() {
	mp := make(map[string]interface{})
	mp["john"] = 100
	mp["dragon"] = 1091
	mp["mother"] = "hero"
	mp["jiongen"] = 12390123.20
	u := User{
		Name:     "mahuateng",
		NickName: "tod",
		Age:      18,
		// Phone:    "110",
	}
	mp["user"] = u
	jsonMarshal(mp)
}

func testStruct() {
	u := User{
		Name:     "mahuateng",
		NickName: "tod",
		Age:      18,
		// Phone:    "110",
	}

	jsonMarshal(u)
}

func jsonMarshal(data interface{}) {
	info, err := json.Marshal(data)

	if err != nil {
		fmt.Printf("json marshal failed, error :%s\n", err)
	}

	fmt.Println(string(info))
}

func testSlice() {
	var s []map[string]interface{}
	mp := make(map[string]interface{})
	mp["john"] = 100
	mp["dragon"] = 1091
	mp["mother"] = "hero"
	mp["jiongen"] = 12390123.20
	u := User{
		Name:     "mahuateng",
		NickName: "tod",
		Age:      18,
		// Phone:    "110",
	}
	mp["user"] = u
	s = append(s, mp)
	mp = make(map[string]interface{})
	mp["john"] = 10011
	mp["dragon"] = 109111
	mp["mother"] = "hero11"
	s = append(s, mp)
	jsonMarshal(s)

}

// json序列化结构体
func main() {
	testStruct()
	// {"Name":"mahuateng","NickName":"tod","Age":18,"Phone":""}
	// 使用json tag
	// {"username":"mahuateng","NickName":"tod","Age":18,"Phone":""}

	testMap()
	// {"dragon":1091,"john":100}

	testSlice()
	// [{"dragon":1091,"jiongen":12390123.2,"john":100,"mother":"hero","user":{"username":"mahuateng","NickName":"tod","Age":18,"Phone":""}},
	// {"dragon":109111,"john":10011,"mother":"hero11"}]
}
