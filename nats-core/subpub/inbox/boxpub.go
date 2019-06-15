package main

import (
	"github.com/nats-io/nats.go"
)

func main(){
	nc,_ := nats.Connect(nats.DefaultURL)

	nc.Subscribe("box", func(msg *nats.Msg) {
		nc.Publish(msg.Reply,[]byte("box i receiver,this is result"))
	})
}