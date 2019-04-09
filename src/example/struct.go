package main

import "fmt"

type student struct {
	Name string
	Age int
	Score int
}
//方法的定义需指定哪个类型的哪个变量，访问控制同样通过大小写控制
func (p *student) init(name string,age int,score int){
	p.Name = name
	p.Age = age
	p.Score = score
}
func main(){
	var stu student
	stu.init("xiaoma",18,99)
	fmt.Println(stu)
}