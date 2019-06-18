### nats-steaming-server usage

usage example:

* [普通的订阅模式](https://github.com/shabbyboy/nats-example/tree/master/nats-stream/basic)

* [队列订阅模式](https://github.com/shabbyboy/nats-example/tree/master/nats-stream/queue)

nats-steaming-server 部署：

* 单机模式：

    > nats-streaming-server -DV

* 集群模式

  参考 [集群配置详细介绍](https://github.com/nats-io/nats-streaming-server#clustering)
  
  首先，nats-streaming-server 是基于nats-server核心开发的，可以理解成nats-server 的上层建筑，所以nats-streaming-server部署分两种情况
  
  nats-server单机 + nats-streaming-server集群
  ```$xslt
  node 1:
  nats-streaming-server -store file -dir store-a -clustered -cluster_node_id a -cluster_peers b,c -nats_server nats://localhost:4222
  node 2:
  nats-streaming-server -store file -dir store-b -clustered -cluster_node_id b -cluster_peers a,c -nats_server nats://localhost:4222
  node 3:
  nats-streaming-server -store file -dir store-c -clustered -cluster_node_id c -cluster_peers a,b -nats_server nats://localhost:4222
  `````
  
  nats-server集群 + nats-streaming-server集群
  
  配置文件：
  
  config node 1
  ```$xslt
    # NATS specific configuration
    port: 4222
    cluster {
      listen: 0.0.0.0:6222
      routes: ["nats://host2:6222", "nats://host3:6222"]
    }
    
    # NATS Streaming specific configuration
    streaming {
      id: mycluster
      store: file
      dir: store
      cluster {
        node_id: "a"
        peers: ["b", "c"]
      }
    }
  ```
  
  config node2
  
  ```$xslt
  # NATS specific configuration
  port: 4222
  cluster {
    listen: 0.0.0.0:6222
    routes: ["nats://host1:6222", "nats://host3:6222"]
  }
  
  # NATS Streaming specific configuration
  streaming {
    id: mycluster
    store: file
    dir: store
    cluster {
      node_id: "b"
      peers: ["a", "c"]
    }
  }
  ```
  
  config node 3
  
  ```$xslt
  # NATS specific configuration
  port: 4222
  cluster {
    listen: 0.0.0.0:6222
    routes: ["nats://host1:6222", "nats://host2:6222"]
  }
  
  # NATS Streaming specific configuration
  streaming {
    id: mycluster
    store: file
    dir: store
    cluster {
      node_id: "c"
      peers: ["a", "b"]
    }
  }
  ```