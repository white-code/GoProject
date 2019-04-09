package main

import (
	"Balance1"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main(){

	var insts []*Balance1.Instance
	for i := 0;i<16;i++{
		host := fmt.Sprintf("192.168.0.%d",rand.Intn(255))
		insts = append(insts,Balance1.NewInstance(host,3320))
	}
	//blance := &Balance1.Randonmethod{}
	//inst,err := blance.Doblance(insts)
	//if err != nil{
	//	fmt.Println(err)
	//}
	var balanceName = "random"
	if len(os.Args) > 1{
		balanceName = os.Args[1]
	}
	Balance1.PrintMethod()
	for {
		inst,err := Balance1.Doblance(balanceName,insts)
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}


}
