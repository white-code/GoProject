package main

import (
	"fmt"
	"strings"
)

//eg1
func Adder() func(int)int{
	var x int
	return func(d int) int{
		x+=d
		return x
	}
}
//闭包实例eg2
func makeSuffix(suffix string) func(string) string{
	return func(name string) string{
		if strings.HasSuffix(name,suffix) == false{
			return name + suffix
		}
		return name
	}
}
//test 数组
func test1(){
	var a[10] int
	b := a
	b[0] = 100
	fmt.Println(a,b)
}

func test2(arr []int){

}
func main(){
	//f:=Adder()
	//fmt.Print(f(100))
	//fmt.Println(f(200))
	//fmt.Println(f(3000))
	//
	////eg2
	//f1 := makeSuffix(".jpg")
	//f2 := makeSuffix(".pdf")
	//
	//fmt.Println(f1("搞笑"))
	//fmt.Println(f2("汽车.pdf"))
	var a[]int
	a = append(a, 10)
	test1()
	test2(a)
	var arr = [...]int{1,2,3}
	arr[1] = 10
	slice := arr[1:2]
	slice[0] = 100
	fmt.Println(arr,"   ",slice)
	slice = append(slice,10)
	fmt.Println(arr,"   ",slice)
	slice = append(slice,10,20,30)
	fmt.Println(arr,"   ",slice)
	s := "dfhajshdf"
	s1 := []rune(s)
	s1[0]='o'

	//sort.Ints()整数排序 数组是值类型，若要排序，需传引用
}