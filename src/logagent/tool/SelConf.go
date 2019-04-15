package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"logagent/tailf"
	"time"
)

const (
	EtcdKey = "/logagent/config/192.168.199.162"
)

//type LogConf struct {
//	Path  string `json:"path""`
//	Topic string `json:"topic"`
//}

func SetLogConfToEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed err :", err)
		return
	}
	fmt.Println("connect succ1")
	defer cli.Close()

	//配置获取日志项
	var logConfArr []tailf.Collectconf
	logConfArr = append(logConfArr, tailf.Collectconf{
		Log_path: "./my.log",
		Topic:    "nginx_log",
	})

	//json打包数据
	data, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("json failed", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//etcd设置日志文件key
	_, err = cli.Put(ctx, EtcdKey, string(data))
	defer cancel()
	if err != nil {
		fmt.Println("Put failed : ", err)
		return
	}
	_, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := cli.Get(ctx, EtcdKey)
	if err != nil {
		fmt.Println("Get failed")
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s,%s \n", ev.Key, ev.Value)
	}
}

func main() {
	SetLogConfToEtcd()

}
