package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main(){
	consumer,err := sarama.NewConsumer(strings.Split("127.0.0.1:9092",","),nil)
	if err != nil{
		fmt.Println("failed to start consumer :",err)
		return
	}
	partitionList,err := consumer.Partitions("nginx_log")
	if err!= nil{
		fmt.Println("Failed to get the list of partitions:",err)
		return
	}
	for partition := range partitionList{
		pc,err := consumer.ConsumePartition("nginx_log",int32(partition),sarama.OffsetNewest)
		if err!= nil{
			fmt.Println("Failed to start the consumer for partition: %d,%s",partition,err)
			return
		}
		defer pc.AsyncClose()
		go func(partitionConsumer sarama.PartitionConsumer){
			wg.Add(1)
			for msg := range pc.Messages(){
				fmt.Printf("partition:%d,Offset:%d,Key:%s,Value:%v\n",msg.Partition,msg.Offset,string(msg.Key),string(msg.Value))
			}
			wg.Done()
		}(pc)
	}
	wg.Wait()
	consumer.Close()

}