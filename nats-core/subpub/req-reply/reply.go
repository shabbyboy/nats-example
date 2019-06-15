package req_reply

import (
	"fmt"
	"github.com/nats-io/go-nats"
)

func main(){

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil{
		fmt.Println("connect error",err)
	}
	//defer nc.Close()
	fmt.Println("start subscribe:")
	nc.Subscribe("help", func(msg *nats.Msg) {
		fmt.Println("received message:",string(msg.Data))
		nc.Publish(msg.Reply,[]byte("i can help you"))
	})

	/// 不是说祝函数调用runtime.Goexit() 会panic吗？
	//go func() {
	//	//var i int = 0
	//	for{
	//		if r := recover(); r != nil{
	//			fmt.Println("panic",r)
	//		}
	//	}
	//}()

	select{

	}
	fmt.Println("go exit run this ")

}
