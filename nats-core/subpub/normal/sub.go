package normal

import (
	"flag"
	"fmt"
	"github.com/nats-io/go-nats"
)

func Subfunc(subj string,url string){
	nc,_ := nats.Connect(url)
	fmt.Println("listen topic:",subj)
	nc.Subscribe(subj, func(msg *nats.Msg) {
		fmt.Println("received a message:",string(msg.Data))
	})
}


func main(){

	var url, subj string
	flag.StringVar(&url,"url",nats.DefaultURL,"nats-server url")
	flag.StringVar(&subj,"subj","","sub topic")

	flag.Parse()

	Subfunc(subj,url)
	//会阻赛 不建议祝函数允许 Goexit() 此函数会立即终止goroutine
	//目前发现好像没发终止main() 函数啊
	//runtime.Goexit()
	select {

	}
}

