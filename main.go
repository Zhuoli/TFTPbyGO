package main 

import (
	"server"
	"flag"
	"os"
	"os/signal"
	"log"
)
var (
	port int
	host string
)

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "Host to use for server")
	flag.IntVar(&port, "port", 0, "Port to use for server")
}

func main() {
	flag.Parse()
	server:=server.NewServer(host,port)	
	// handle ctrl-c
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		//range iterates over of elements in a variety of data structures.
		for sig := range c {
			log.Printf("Received %v, exiting", sig)
			server.ServerConn.Close()
			os.Exit(0)
		}
	}()
	server.Run()
}