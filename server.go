
package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"io"
	"log"
	"net/http"
)


func server(port *int) {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "OK!")
	})
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Starting http server %s", addr)
	var c, _, err = zk.Connect(servers, zk_connect_timeout)
	defer c.Close()
	if err != nil {
		panic(err)
	}

	initServerNode(c)

	if err := RegistServer(c, port); err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(addr, nil);err == nil {
		log.Fatal(err)
	}
}

func RegistServer(c *zk.Conn, p *int) error {
	paths := server_node+"/"+fmt.Sprintf("%d", *p)
	flg, _, err := c.Exists(paths)
	if err != nil {
		return err
	}
	if flg {
		log.Printf("Port %s has been startted.", *p)
		return nil
	}

	s, err := c.Create(paths, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err == nil {
		fmt.Println(s)
		return nil
	} else {
		return  err
	}
}