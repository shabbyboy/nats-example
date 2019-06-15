### nats-server usage example

* example
    * 普通的发布订阅 example [link](https://github.com/shabbyboy/nats-example/tree/master/nats-core/subpub/normal)
    * queue group example [link](https://github.com/shabbyboy/nats-example/tree/master/nats-core/subpub/queue)
    * request/reply example [link](https://github.com/shabbyboy/nats-example/tree/master/nats-core/subpub/req-reply)
    * chananel example [link](https://github.com/shabbyboy/nats-example/tree/master/nats-core/subpub/chan)
    * encode example [link](https://github.com/shabbyboy/nats-example/tree/master/nats-core/subpub/encode)
    * context example [link](https://github.com/shabbyboy/nats-example/tree/master/nats-core/subpub/context)
    * inbox example [link](https://github.com/shabbyboy/nats-example/tree/master/nats-core/subpub/inbox)


* nats-server configure

   * 单机部署
        > nats-server -DV
        
   * 集群部署
   
       ```$xslt
       node one：
       nats-server -p 4222 -m 4333 -cluster nats://localhost:4248 -routes nats://localhost:5248,nats://localhost:6248 -DV
       node two：
       nats-server -p 5222 -m 5333 -cluster nats://localhost:5248 -routes nats://localhost:4248,nats://localhost:6248 -DV
       node three：
       nats-server -p 6222 -m 6333 -cluster nats://localhost:6248 -routes nats://localhost:5248,nats://localhost:4248 -DV
        ```
        
        * -p ：指定nats 连接端口
        * -m ：指定监视服务端口
        * -cluster ：指定本节点在集群内暴露地址，单个集群的节点通过路由通信
        * -routes ：指定本节点路由到其他节点的地址
        * -DV ：在终端显示日志、stack信息