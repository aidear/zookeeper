package main

import "time"

var servers  = []string{"127.0.0.1:2181","127.0.0.1:2182","127.0.0.1:2183"}
var (
	h bool
	m bool
	s bool
	port *int
)

const (
	server_node = "/online_servers"
	zk_connect_timeout = time.Second*10
)