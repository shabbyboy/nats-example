package main

import (
	"github.com/nats-io/nats.go"
)

func main(){
	nc, _ := nats.Connect(nats.DefaultURL)

	//note: encode 模式也是支持队列模式的，这里就不写demo了

	nc.Publish("queue",[]byte("nihao queue"))

	defer nc.Close()

}