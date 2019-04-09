package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/wendal/errors"
	"logagent/tailf"
)

var (appconfig *Config)

type Config struct{
	log_level string
	log_path string
	kafka_addr string
	collectconf []tailf.Collectconf
	etcdAddr string
	etcdKey string
}



func loadCollectConf(conf config.Configer)(cc tailf.Collectconf ,err error){
	cc.Log_path = conf.String("collect::log_path")
	if len(cc.Log_path) == 0{
		err = errors.New("invalid collect::log_path")
	}
	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0{
		err = errors.New("invalid collect::topic")
	}
	return
}

func loadconf(conftype,filename string)(err error){
	conf,err := config.NewConfig(conftype,filename)
	if err != nil{
		fmt.Println("NewConfig failed err:",err)
		return
	}
	//初始化配置结构体，
	appconfig = &Config{}
	//获取打印日志级别和打印路径
	appconfig.log_level = conf.String("logs::log_level")
	if len(appconfig.log_level) == 0{
		fmt.Println("invild log_level")
		appconfig.log_level = "debug"
	}
	appconfig.log_path = conf.String("logs::log_path")
	if len(appconfig.log_level) == 0{
		fmt.Println("invild log_path")
		appconfig.log_path = "./logs"
	}
	//获取kafka地址
	appconfig.kafka_addr = conf.String("kafka::server_addr")
	if len(appconfig.kafka_addr) == 0{
		fmt.Println("invild kafka_addr")
		return
	}
	//获取收集日志路径和topic
	cc,err := loadCollectConf(conf)
	if err != nil{
		fmt.Println("loadCollectConf failed err:",err)
		return
	}
	appconfig.collectconf = append(appconfig.collectconf,cc)

	//获取etcd相关信息
	appconfig.etcdAddr = conf.String("etcd::addr")

	if len(appconfig.etcdAddr) == 0{
		fmt.Println("invild etcdAddr")
		return
	}
	appconfig.etcdKey = conf.String("etcd::configKey")

	if len(appconfig.etcdKey) == 0{
		fmt.Println("invild etcdKey")
		return
	}

	return
}