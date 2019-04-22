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

func testStruct() (ret string) {
	u := User{
		Name:     "马化腾",
		NickName: "tod",
		Age:      18,
		// Phone:    "110",
	}

	ret, _ = jsonMarshal(u)
	return
}

func jsonMarshal(data interface{}) (ret string, err error) {
	info, err := json.Marshal(data)

	if err != nil {
		fmt.Printf("json marshal failed, error :%s\n", err)
		return
	}

	fmt.Println(string(info))
	ret = string(info)
	return
}

func testMap() (ret string, err error) {
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
	ret, err = jsonMarshal(mp)
	return
}

func testUnmarshalStruct() {
	var user User
	userInfo := testStruct()
	err := json.Unmarshal([]byte(userInfo), &user)
	if err != nil {
		fmt.Errorf("Json Unmarshal failed, err: %s\n", err)
	}
	fmt.Println(user)
	// {"username":"马化腾","NickName":"tod","Age":18,"Phone":""}
	// {马化腾 tod 18 }
}

func testUnmarshalMap() {
	var mp map[string]interface{}
	// mp = make(map[string]interface{})
	jsonStr, _ := testMap()
	err := json.Unmarshal([]byte(jsonStr), &mp) // 在内部分配内存，所以需要传入地址
	if err != nil {
		fmt.Errorf("Json Unmarshal failed, err: %s\n", err)
	}
	fmt.Println(mp)
}

func main() {
	testUnmarshalStruct()
	testUnmarshalMap()
}
