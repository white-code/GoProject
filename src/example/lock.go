package main

import (
	"fmt"
	"sync"
)

//互斥锁
var lock sync.Mutex
func testarr(a *[4]int){
	(*a)[0] = 100
}

func testlock(){
	arr := [...]int{1,2,3,4}
	a := make(map[int]int)
	a[0] = 100
	a[2] = 8
	for i := 0;i<2; i++{
		lock.Lock()
		go func(arr map[int]int) {
			arr[0] = 1000
		}(a)
		lock.Unlock()
	}
	//lock.Lock()
	//testarr(&arr)
	//arr[0] = 100
	fmt.Println(a)
	fmt.Println(arr)
	//lock.Unlock()

}
func main(){
	testlock()
}