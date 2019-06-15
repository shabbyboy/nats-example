package encode

import (
	"fmt"
	"github.com/nats-io/go-nats"
)

type person struct {
	Name string
	Address string
	age int
}

func main(){

	nc, _ := nats.Connect(nats.DefaultURL)

	c, _ := nats.NewEncodedConn(nc,nats.JSON_ENCODER)

	defer c.Close()

	/*
	nats.JSON_ENCODER 数据格式是json 所以结构体解析需要满足go json 解析的标准，比如属性需要是大写属性
	 */
	c.Subscribe("foo", func(s string) {
		fmt.Println("json receive:",s)
	})

	c.Subscribe("person", func(p *person) {
		fmt.Println("received person:",p)
	})

	// func 参数：
	// 1个：msg.data
	// 2个：msg.subject msg.data
	// 3个：msg.subject msg.reply msg.data
	c.Subscribe("help", func(subj string,reply string,msg string) {
		fmt.Println("receive below message:",subj,msg)
		c.Publish(reply,"ican help")
	})


	select{

	}

}