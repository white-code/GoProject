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
		fmt.Println("connect failed err :",err)
		return
	}
	fmt.Println("connect succ")
	defer cli.Close()
	_,cancel := context.WithTimeout(context.Background(),5*time.Second)
	_,err = cli.Put(context.Background(),"./logcollect","sample_value")
	defer cancel()
	if err!=nil{
		fmt.Println("Put failed : ",err)
		return
	}
	_,cancel = context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	resp,err := cli.Get(context.Background(),"./logcollect")
	if err!=nil{
		fmt.Println("Get failed")
		return
	}
	for _,ev := range resp.Kvs{
		fmt.Printf("%s,%s \n",ev.Key,ev.Value)
	}

}
