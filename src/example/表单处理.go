package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const form = `<html><body><form action ="#" method="post" name="bar"><input type="text" name = "in"/><input type="text" name = "in"/><input type="submit" value="submit"/></form></body></html>`
func Simpleserve(w http.ResponseWriter,request *http.Request){
	panic("test")
	io.WriteString(w,"<h1>hello,world</h1>")
}

func FormServe(w http.ResponseWriter,request *http.Request){
	w.Header().Set("Content-Type","text/html")//设置头部
	switch request.Method {
	case "GET":
		io.WriteString(w,form)
	case "POST":
		request.ParseForm()//解析表单数据
		io.WriteString(w,request.Form["in"][1])
		io.WriteString(w,"\n")
		io.WriteString(w,request.FormValue("in"))
	}
}

func logPanic(handle http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func(){
			if x := recover(); x != nil{
				log.Printf("[%v] caught panic:%v",request.RemoteAddr,x)
			}
		}()
		handle(writer,request)
	}
}
func main(){
	http.HandleFunc("/test1",logPanic(Simpleserve))
	http.HandleFunc("/test2",logPanic(FormServe))
	if err := http.ListenAndServe(":8088",nil);err != nil{
		fmt.Println(err)
	}
}