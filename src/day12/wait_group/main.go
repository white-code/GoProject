package main


import (
	"fmt"
	"sync"
	"time"
)
func main(){
	wg := sync.WaitGroup{}
	wg.Add(10) //防止goroutine 尚未启动主进程就结束
	for i := 0 ; i<10 ; i++{
		go calc(&wg,i)
	}
	wg.Wait()
	fmt.Println("all goroutine finish")
}


func calc(w *sync.WaitGroup,i int){
	//w.Add(1)
	fmt.Println("calc:",i)
	time.Sleep(time.Second)
	w.Done()
}