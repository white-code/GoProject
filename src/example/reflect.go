package main

import (
	"fmt"
	"reflect"
)
type Student struct {
	Name string
	Age int
}

func text(b interface{}){
	v := reflect.TypeOf(b)
	fmt.Println(v)
	t := reflect.ValueOf(b)
	k := t.Kind()
	fmt.Println(t,k)

}

func main(){
	var a Student = Student{Name:"malinchao",Age:22}
	text(a)
}