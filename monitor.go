package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
)

func monitor() {
	var c, _, err = zk.Connect(servers, zk_connect_timeout)

	if err != nil {
		panic(err)
	}
	defer c.Close()

	for {
		_, _, ch, err := c.ChildrenW(server_node)
		if err != nil {
			panic(err)
		}

		if event, ok := <-ch;ok {
			fmt.Println(event.Type)
			if event.Type == zk.EventNodeChildrenChanged {
				getOnlineServer(c, server_node)
			}
		}
		//time.Sleep(time.Millisecond * 10)
	}
}

func getOnlineServer(c *zk.Conn, paths string) {
	if hosts, _, err := c.Children(paths); err == nil {
		fmt.Println("在线server列表：")
		//if len(hosts) <= 0 {
		//	fmt.Println("暂无在线server。")
		//}
		for _, v := range hosts {
			fmt.Printf(":%s\n",v)
		}
	}
}