package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string
	Passward string
	Age int
	Sex string
}

func testStruct(){
	usr := &User{"malinchao","wkeewj123",18,"man"}
	data,err := json.Marshal(usr)
	if err != nil{
		fmt.Printf("json.Mashal failed err: ",err)
		return
	}
	fmt.Println(string(data))
	return
}

func testMap1(){
	m := make(map[string]interface{})
	m["user"] = "djhf"
	m["age"] = 18
	data,err := json.Marshal(m)
	if err != nil{
		fmt.Printf("json.Mashal failed err: ",err)
		return
	}
	fmt.Println(string(data))
	return
}
func testSlice(){
	var s []interface{}
	m := make(map[string]interface{})
	m["user"] = "djhf"
	m["age"] = 18
	s = append(s, m)
	l := make(map[string]interface{})
	l["user"] = "dafsd"
	l["age"] = 19
	s = append(s, l)
	data,err := json.Marshal(s)
	if err != nil{
		fmt.Printf("json.Mashal failed err: ",err)
		return
	}
	fmt.Println(string(data))
	return
}
func main(){
	testStruct()
	testMap1()
	testSlice()
}
