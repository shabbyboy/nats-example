package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main(){
	nc,_ := nats.Connect(nats.DefaultURL)

	//box 的用法和request reply 相同类似，requst 不需要另起一个订阅来接受消息
	inbox := nats.NewInbox()

	sub,err := nc.Subscribe(inbox, func(msg *nats.Msg) {
		fmt.Println("inbox reciver:")
		fmt.Println(msg)
	})

	if err !=nil{
		return
	}

	defer func() {
		nc.Close()
		sub.Unsubscribe()
	}()

	msg := &nats.Msg{
		Subject: "box",
		Reply:   inbox,
		Data:    []byte("nihao box"),
	}
	fmt.Println("pub inbox ",msg)
	nc.PublishMsg(msg)

	select{

	}

}