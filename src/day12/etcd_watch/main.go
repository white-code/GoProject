package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main(){
	cli,err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"localhost:2379"},
		DialTimeout:5*time.Second,
	})
	if err != nil{
		fmt.Println("connect failed err: ",err)
		return
	}
	fmt.Println("connect succ")
	_,err = cli.Put(context.Background(),"./logcollect","23456")
	if err!=nil{
		fmt.Println("Put failed : ",err)
		return
	}

	defer cli.Close()
	for {
		rch :=cli.Watch(context.Background(),"./logcollect")
		for wresp := range rch{
			for _,ev := range wresp.Events{
				fmt.Println(ev.Type,ev.Kv.Key,ev.Kv.Value)
			}
		}
	}
}
