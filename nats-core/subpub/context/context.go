package main

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

func main(){

	//context 模式官方文档并没有介绍，这种用法可以用来统一管理订阅，比如：函数结束，取消订阅
	ctx,cancel := context.WithTimeout(context.Background(),2*time.Second)

	defer cancel()

	nc, _ := nats.Connect(nats.DefaultURL)

	msg, _ := nc.RequestWithContext(ctx,"foo",[]byte("bar"))

	sub, _ := nc.SubscribeSync("foo")
	msg2, _ := sub.NextMsgWithContext(ctx)

	// Encoded Request with context
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	fmt.Println(msg,msg2,c)


}
