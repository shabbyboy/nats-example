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
	defer cc.Close()

	recvCh := make(chan *person)

	cc.BindRecvChan("hello.*",recvCh)
	fmt.Println("start recive msg")
	select{
	case who := <- recvCh:
		fmt.Println("msg:")
		fmt.Println(who)
	}
}

