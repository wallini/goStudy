package main

import "fmt"

import "encoding/json"

func main() {
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
