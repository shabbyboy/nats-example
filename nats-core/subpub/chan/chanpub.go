package _chan

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main(){

	nc, _ := nats.Connect(nats.DefaultURL)

	cc, _ := nats.NewEncodedConn(nc,nats.JSON_ENCODER)

	type person struct{
		Name string
		Address string
		Age int
	}

	sendCh := make(chan *person)
	//bindSendChan 原理是开启了一个goroutine 持续接收sendCh，并发送，
	// 所以实际使用，在祝函数需要阻赛，否则主函数会立马结束无法完成消息的传递
	//note： 这里发现chanel 是通过反射实现的，不知道性能怎么样
	cc.BindSendChan("hello.me",sendCh)

	me := &person{Name:"derek",Age:22,Address:"140 new montgomery street"}
	fmt.Println("start pub msg:",me)
	defer cc.Close()
	sendCh <- me

	select {

	}
}
