package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"logagent/kafka"
	"logagent/tailf"
)

func main() {
	file_name := "./logcollect.conf"
	err := loadconf("ini", file_name)
	if err != nil {
		fmt.Printf("load conf failed,err:%v", err)
		return
	}
	fmt.Println(appconfig.log_path, appconfig.log_level, appconfig.collectconf)
	err = initLogger()
	if err != nil {
		fmt.Printf("loadLogs() failed,err:%v", err)
		return
	}
	collectConf, err := initEtcd(appconfig.etcdAddr, appconfig.etcdKey)
	if err != nil {
		logs.Error("etcd init failed,err:%v", err)
		return
	}
	err = tailf.Inittail(collectConf)
	if err != nil {
		fmt.Printf("inittail() failed,err:%v", err)
		return
	}
	err = kafka.InitKafka(appconfig.kafka_addr)
	if err != nil {
		fmt.Printf("InitKafka() failed,err:%v", err)
		return
	}
	fmt.Println("log system init success")

	err = serverrun()
	if err != nil {
		logs.Error("serverrun failed,err %v", err)
		return
	}

}
