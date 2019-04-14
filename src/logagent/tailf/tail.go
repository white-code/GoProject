package tailf

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"time"
)

type TextMsg struct {
	Msg   string
	Topic string
}

type Collectconf struct {
	Log_path string `json:"log_path"`
	Topic    string `json:"topic"`
}

type TailObj struct {
	tail *tail.Tail
	conf Collectconf
}

type TailObjMgr struct {
	TailObjs []*TailObj
	msgChan  chan *TextMsg
}

var (
	tailObjMgr *TailObjMgr
)

func GetOneLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return
}

func readFormTail(tailObj *TailObj) {
	var line *tail.Line
	var ok bool
	for true {
		line, ok = <-tailObj.tail.Lines
		if !ok {
			logs.Warn("tail file close reopen,filename :%s\n", tailObj.tail.Filename)
			time.Sleep(100 * time.Microsecond)
			continue
		}
		textMsg := &TextMsg{
			Msg:   line.Text,
			Topic: tailObj.conf.Topic,
		}
		tailObjMgr.msgChan <- textMsg
	}
}

func Inittail(collectconf []Collectconf) (err error) {
	if len(collectconf) == 0 {
		err = fmt.Errorf("incalid config for log collect,conf:%v", collectconf)
		return
	}
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, 100),
	}
	for _, v := range collectconf {
		obj := &TailObj{conf: v}
		filename := v.Log_path
		fmt.Println(filename)
		tails, Tailerr := tail.TailFile(filename, tail.Config{
			ReOpen:    true,
			Follow:    true,
			Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
			MustExist: false,
			Poll:      true,
		})
		obj.tail = tails
		if Tailerr != nil {
			fmt.Println("TailFile 出错:", err)
			return
		}

		tailObjMgr.TailObjs = append(tailObjMgr.TailObjs, obj)
		go readFormTail(obj)

	}

	return
}
