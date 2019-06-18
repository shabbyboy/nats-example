package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

func main(){

	/*
	note 每一个客户端连接 客户端id 必须是唯一的，sub： clinetone  pub： clienttwo
	 */
	sc, _ := stan.Connect("test-cluster","clienttwo")

	fmt.Println("publish a message:")
	//msg := stan.Msg{
	//
	//}
	sc.Publish("foo",[]byte("hello world4"))

	defer sc.Close()
}