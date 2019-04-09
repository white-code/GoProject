package add

import "fmt"

var Name string //= "malinchao"
var Age int = 1000 //= 24小写开头为私有，在别的包无法调用

func init() {
	fmt.Println("this is add init")
	Name = "malinchao"
	Age = 18
}
func add1(arg...int) int {
	n:=len(arg)
	var sum int = 0
	for i:=0 ;i<n;i++ {
		sum+=arg[i]
	}
	return sum
}