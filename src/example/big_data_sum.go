package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func strAdd(str1,str2 string) (result string) {
	if len(str1)==0&&len(str2)==0{
		result = "0"
		return
	}
	index1 := len(str1)-1
	index2 := len(str2)-1
	var flag int = 0
	for index1>=0 && index2>=0{
		c1 := str1[index1] - '0'
		c2 := str2[index2] - '0'
		sum := int(c1)+int(c2)+flag
		if sum >= 10{
			flag = 1
		}else {
			flag = 0
		}
		c3 := (sum%10) + '0'
		result = fmt.Sprintf("%c%s",c3,result)
		index1--
		index2--
	}
	for index1 >= 0{
		c1 := str1[index1] - '0'
		sum := int(c1) + flag
		if sum >= 10{
			flag = 1
		}else {
			flag = 0
		}
		c3 := (sum%10) + '0'
		result = fmt.Sprintf("%c%s",c3,result)
		index1--
	}
	for index2 >= 0{
		c1 := str2[index2] - '0'
		sum := int(c1) + flag
		if sum >= 10{
			flag = 1
		}else {
			flag = 0
		}
		c3 := (sum%10) + '0'
		result = fmt.Sprintf("%c%s",c3,result)
		index2--
	}
	if flag == 1{
		result = fmt.Sprintf("1%s",result)
	}
	return
}
func main()  {
	reader := bufio.NewReader(os.Stdin)
	result,_,err := reader.ReadLine()
	if err!=nil{
		fmt.Println("读取出错：",err)
		return
	}
	strSlice := strings.Split(string(result),"+")
	if len(strSlice) != 2 {
		fmt.Println("请输入a+b格式")
		return
	}
	strNumber1 := strings.TrimSpace(strSlice[0])
	strNumber2 := strings.TrimSpace(strSlice[1])

	fmt.Println(strAdd(strNumber1,strNumber2))

}