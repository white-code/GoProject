package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (client sarama.SyncProducer)

//初始化kafka
func InitKafka(adr string)(err error){
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a test message log")

	client,err = sarama.NewSyncProducer([]string{adr},config)

	if err != nil{
		logs.Error("init kafka producer failed,err:%v",err)
		return
	}
	logs.Debug("init kafka success")
	return
}

//向kafka发送信息
func SendToKafka(data,topic string)(err error){
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	pid,offset,Serr := client.SendMessage(msg)

	if err != nil{
		err = Serr
		logs.Error("send message failed, err: %v,topic: %s,data: %s",err,topic,data)
		return
	}
	logs.Debug("send succ,pid: %v,offset: %v,topic: %v",pid,offset,topic)
	return
}