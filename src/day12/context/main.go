package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct{
	r *http.Response
	err error
}
func process(){
	ctx,cancel := context.WithTimeout(context.Background(),2*time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport:tr}
	c := make(chan *Result,1)
	req,err := http.NewRequest("GET","http://www.baidu.com",nil)
	if err != nil{
		fmt.Println("http request failed,err: ",err)
		return
	}
	go func(){
		resp,err := client.Do(req)
		pack := &Result{resp,err}
		c <- pack
		return
	}()
	select {
	case <- ctx.Done()://超时
		tr.CancelRequest(req) //取消返回错误
		<-c//将错误取出
		fmt.Println("time out")
	case res := <-c:
		defer res.r.Body.Close()
		out,_ := ioutil.ReadAll(res.r.Body)
		fmt.Println(out)
	}
}

func main(){
	process()
}
