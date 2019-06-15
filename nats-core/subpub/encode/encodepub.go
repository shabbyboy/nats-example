package encode

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"time"
)


type Person struct {
	Name string
	Address string
	Age int
}

func main(){

	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	// Simple Publisher
	c.Publish("foo", "Hello World")

	me := &Person{Name: "derek", Age: 22, Address: "140 New Montgomery Street, San Francisco, CA"}

	c.Publish("person",me)
	var response string
	err := c.Request("help","hepl me",&response,10*time.Millisecond)

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(response)
}
