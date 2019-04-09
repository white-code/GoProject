package main

import(
	"fmt"
	"github.com/astaxie/beego/config"
	)



func main(){
	conf,err := config.NewConfig("ini","./logcollect.conf")
	if err != nil{
		fmt.Println("NewConfig failed err:",err)
		return
	}
	port,err := conf.Int("server::listen_port")
	if err != nil{
		fmt.Println("read server::port failed,err:",err)
		return
	}
	fmt.Println("Port:",port)
	loglevel := conf.String("logs::log_level")
	if err != nil{
		fmt.Println("read logs::log_level failed,err:",err)
		return
	}
	fmt.Println("log_level:",loglevel)
	logPath := conf.String("logs::log_path")
	fmt.Println("log_path:", logPath)
}
