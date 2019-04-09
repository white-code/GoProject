package main

import (
	"fmt"
	"net"
)

//项目上线，资源管理系统需要给项目分配服务器，公司中通过项目ID去资源管理系统中获取项目对应的ip


var (
	localIpArray []string
)

func init(){
	addrs,err := net.InterfaceAddrs()
	if err != nil{
		panic(fmt.Sprint("get local ip failed"))
	}
	for _,addr := range addrs{
		if ipnet,ok := addr.(*net.IPNet);ok && !ipnet.IP.IsLoopback(){
			if ipnet.IP.To4() != nil{
				localIpArray = append(localIpArray,ipnet.IP.String())
			}
		}
	}
	fmt.Println(localIpArray)
}