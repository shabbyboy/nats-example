package main

import "github.com/nats-io/stan.go"

func main(){

	sc, _ := stan.Connect("test-cluster","queuetwo")

	sc.Publish("tfoo",[]byte("nihao wangba duzi"))

	sc.Close()

}
