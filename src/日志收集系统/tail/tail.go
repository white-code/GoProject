package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)


func main(){
	filename := "./my.log"
	tails,err := tail.TailFile(filename,tail.Config{
		ReOpen:true,
		Follow:true,
		Location:&tail.SeekInfo{Offset:0,Whence:2},
		MustExist:false,
		Poll:true,
	})
	if err != nil{
		fmt.Println("TailFile 出错:", err)
		return
	}
	var msg *tail.Line
	var ok bool
	for true{
		msg,ok = <- tails.Lines
		if !ok{
			fmt.Println("tail file close reopen :%s\n",tails.Filename)
			time.Sleep(100*time.Microsecond)
			continue
		}
		fmt.Println("msg:",msg)
	}
}
