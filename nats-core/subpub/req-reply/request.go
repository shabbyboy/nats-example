package req_reply

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"time"
)

func main(){

	nc, _ := nats.Connect(nats.DefaultURL)

	msg, err := nc.Request("help", []byte("help me"),10*time.Second)

	if err != nil{
		fmt.Println("request error",err)
	}

	fmt.Println(string(msg.Data))
}
