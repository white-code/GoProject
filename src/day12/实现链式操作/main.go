package main

import "fmt"

type stu struct{
	Name string
	Age int
}

func(p *stu) SetAge(n int)(q *stu){
	p.Age = n
	q = p
	return
}

func(p *stu) SetName(name string)(q *stu){
	p.Name = name
	q = p
	return
}


func main(){
	var n = int(18)
	var name = string("malinchao")
	var p stu
	p.SetAge(n).SetName(name)
	fmt.Println(p)
}