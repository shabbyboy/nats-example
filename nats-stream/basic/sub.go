package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

func main(){

	sc, _ := stan.Connect("test-cluster","clientone")

	// nats-streaming-server 不支持 通配符 主题
	/*
	由于持久化的特性，nats-streaming-server 是可以订阅到在客户端注册之前就发布的信息的，可以通过如下设置：
	stan.StartWithLastReceived() //从上次接收的位置开始
	stan.DeliverAllAvailable()  // 所有可以得到的消息
	stan.StartAtSequence(22)   //根据消息的Sequence 来获取
	stan.StartAtTime(startTime) //接收消息的时间戳
	stan.StartAtTimeDelta(time.ParseDuration("30s")）//时间差
	stan.AckWait(20*time.Second) //设置消息未收到重传的延迟

	stan.Msg:{
	Sequence  => 序列号 stan.StartAtSequence
	Subject     订阅主题
	Reply      	request/reply
	Data        message
	Timestamp   消息发送的时间
	Redelivered bool 是否重传的数据
	CRC32
	}
	 */


	sub,_ := sc.Subscribe("foo", func(msg *stan.Msg) {
		fmt.Println("received message",msg.Sequence,msg.Timestamp)
	},stan.DeliverAllAvailable())

	sub.Unsubscribe()

	defer func() {
		// nats-streaming-server 的订阅主题是会存储在服务器端的，所以如果这个订阅后面不再需要的话建议 调用unsubscribe()取消订阅
		sub.Unsubscribe()
		sc.Close()
	}()
	select {

	}
}