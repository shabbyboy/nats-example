package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

func main(){


	/*
	Durable Queue 和 queue，durable 可以在订阅取消，并在此重连之后，接收将从上次退出的位置开始
	note: 具体效果我没测出来 手动🐶
	 */
	sc, _ := stan.Connect("test-cluster","queueone")

	//qsub,_ := sc.QueueSubscribe("tfoo","queue", func(msg *stan.Msg) {
	//	fmt.Println("receive msg:",string(msg.Data),msg.Sequence)
	//})
	//
	//qsubt,_ := sc.QueueSubscribe("tfoo","queue", func(msg *stan.Msg) {
	//	fmt.Println("receive msgtwo:",string(msg.Data),msg.Sequence)
	//})


	qsub,_ := sc.QueueSubscribe("tfoo","queue", func(msg *stan.Msg) {
		fmt.Println("receive msg:",string(msg.Data),msg.Sequence)
	},stan.DurableName("q"))

	qsubt,_ := sc.QueueSubscribe("tfoo","queue", func(msg *stan.Msg) {
		fmt.Println("receive msgtwo:",string(msg.Data),msg.Sequence)
	},stan.DurableName("q"))

	select {

	}

	defer func() {
		qsub.Unsubscribe()
		qsubt.Unsubscribe()
		sc.Close()
	}()

}
