package normal

import (
	"flag"
	"fmt"
	"github.com/nats-io/go-nats"
)

func main(){

	var topic, msg ,url string
	flag.StringVar(&url, "url",nats.DefaultURL,"nats-server url")
	flag.StringVar(&topic,"topic","","pub topic")
	flag.StringVar(&msg,"msg","","deliver msg")

	flag.Parse()

	nc,_ := nats.Connect(url)
	fmt.Println("publisher topic:",topic)
	nc.Publish(topic,[]byte(msg))


}

