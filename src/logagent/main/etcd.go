package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/clientv3"
	"logagent/tailf"
	"strings"
	"time"
)

type EtcdClient struct {
	client *clientv3.Client
}

var (
	etcdClient *EtcdClient
)

func initEtcd(addr string, key string) (collectConf []tailf.Collectconf, err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed:", err)
	}
	etcdClient = &EtcdClient{cli}

	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}
	//var collectConf []tailf.Collectconf
	//从etcd获取配置
	for _, ip := range localIpArray {
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		resp, err := cli.Get(ctx, etcdKey)
		cancel()
		if err != nil {
			logs.Debug("Get failed err :", err)
			continue
		}
		logs.Debug("resp from etcd :%v", resp.Kvs)
		for _, v := range resp.Kvs {
			//fmt.Println(k,v)
			if string(v.Key) == etcdKey {
				err = json.Unmarshal(v.Value, &collectConf) //将获取到的配置返回出去
				if err != nil {
					logs.Error("Unmarshal failed,err :%v", err)
					continue
				}
				fmt.Println(collectConf)
				logs.Debug("log congig is %v", collectConf)
			}

		}

	}

	defer cli.Close()
	return
}
