package main

import (
	"encoding/json"
	"fmt"
)

func test() {
	/*    Name string `json:"msg_name"`       // 对应JSON的msg_name
	Body string `json:"body,omitempty"` // 如果为空置则忽略字段
	Time int64  `json:"-"`              // 直接忽略字段
	*/
	str := `{
		"a" : "44",
		"b" : "heh",
		"c" : 123
	}`
	type Info struct {
		A string `json:"a"`
		B string `json:"b"`
		C int    `json:"c"`
	}
	var myInfo Info
	err := json.Unmarshal([]byte(str), &myInfo)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", myInfo)

	b, err := json.Marshal(myInfo)
	if err != nil {
		fmt.Println("error1:", err)
	}
	fmt.Println(string(b[:]))

}
