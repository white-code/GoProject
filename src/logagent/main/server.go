package main

import (
	"github.com/astaxie/beego/logs"
	"logagent/kafka"
	"logagent/tailf"
	"time"
)

func serverrun()(err error){
	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil{
			logs.Error("send to kafka failed,err: %v",err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg)(err error){
	err = kafka.SendToKafka(msg.Msg,msg.Topic)
	if err != nil{
		logs.Error("SendToKafka err,err: %v",err)
		return
	}
	//logs.Debug("send masg test: msg: %s,topic: %s\n",msg.Msg,msg.Topic)
	return
}