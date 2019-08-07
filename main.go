package main

import (
	"flag"
	"fmt"
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&m, "m", false, "Monitor")
	flag.BoolVar(&s, "s", false, "Start a server")
	port = flag.Int("p", 0, "port: must")
}

func main() {
	//var port *int
	//port = flag.Int("p", 0, "port: must")
	flag.Parse()
	if h {
		Usage()
	}
	if m {
		monitor()
	}

	if s {

		if *port<=0 || *port>65535 {
			flag.Usage()
			return
		}
		server(port)
	}


}


func Usage()  {
	fmt.Println(".")
}