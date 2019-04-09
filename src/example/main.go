package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func urlProcess(url string) string{
	result := strings.HasPrefix(url,"http://")
	if(!result){
		url = fmt.Sprintf("http://%s",url)
	}
	return url
}

func pathProcess(path string) string{
	result := strings.HasSuffix(path,"/")
	if(!result){
		path = fmt.Sprintf("%s/",path)
	}
	return path
}

func main(){
	//var(
	//	url string
	//	path string
	//)
	//fmt.Scanf("%s%s",&url,&path)
	//url = urlProcess(url)
	//path = pathProcess(path)
	//fmt.Printf("url:%s\npath:%s",url,path)
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04"))
	var i int
	fmt.Printf("%d",&i)
	n := rand.Intn(100)

	for{
		var input int
		fmt.Scanf("%d",&input)
		switch {
		case input == n:
				fmt.Println("you are right!!!")
				break
		case input > n:
			fmt.Println("bigger!!!")
		case input < n :
			fmt.Println("less!!!")
		}
	}

}