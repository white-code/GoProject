package main

import "fmt"


type test interface {
	print()
}

type Student struct{
	Name string
	Age int
	Score int
}



func (p *Student) print(){
	fmt.Printf("name: %s\nage:%d\nscore:%d\n",p.Name,p.Age,p.Score)
}
func main(){
	var t test
	stu := Student{Name: "malinchao",Age: 18,Score:99}
	t = &stu
	t.print()
	stu.print()
}