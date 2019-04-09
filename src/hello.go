package main

import (
	"add"
	"fmt" //导入的包不想用，只用来做初始化可用下划线做别名
)

const (
	Man = 1
	Female = 2
)

func list(n int){
	for i:= 0 ; i <= n; i++{
		fmt.Printf("%d+%d=%d\n",i,n-i,n)
	}
}

func isSexUdf(secoond int){
	if (secoond % Female == 0){
		fmt.Println("女生")
	}else {
		fmt.Println("男生")
	}
}
func swqp(a *int, b *int){
	*a = *a+*b
	*b = *a-*b
	*a = *a-*b
}
func reverse(str string) string{
	var result string
	length := len(str)
	for i:=0 ; i<length ; i++{
		result = result + fmt.Sprintf("%c",str[length-i-1])
	}
	return result
}

func main() {
	//list(5)
	fmt.Println(add.Age,add.Name)
	//i := time.Now().Unix()
	//isSexUdf(int(i))
	//goods := os.Getenv("GOOS")
	//fmt.Printf("The operating system is :%s\n",goods)
	//path := os.Getenv("PATH")
	//fmt.Printf("PATH is %s/n",path)
	var a = 10
	var b = 20
	swqp(&a,&b)
	fmt.Printf("a:%d,b:%d",a,b)
	fmt.Println(reverse("adhfahf"))
	sum:=add.add1(10) 
	fmt.Println(sum)
}