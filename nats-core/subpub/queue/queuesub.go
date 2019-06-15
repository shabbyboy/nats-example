package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main(){
	nc, _ := nats.Connect(nats.DefaultURL)

	//note: encode 模式也是支持队列模式的，这里就不写demo了
	//具有同一个queue 组的，只有个一个会处理消息，这可以用来实现负载均衡和容错性

	nc.QueueSubscribe("queue","work", func(msg *nats.Msg) {
		fmt.Println("queue 1 receiver:")
		fmt.Println(string(msg.Data))
	})

	nc.QueueSubscribe("queue","work", func(msg *nats.Msg) {
		fmt.Println("queue 2 receiver:")
		fmt.Println(string(msg.Data))
	})

	nc.QueueSubscribe("queue","work", func(msg *nats.Msg) {

		fmt.Println("queue 3 receiver:")
		fmt.Println(string(msg.Data))
	})

	defer nc.Close()
	select {

	}
}
