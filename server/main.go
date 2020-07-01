package main

import (
	"flag"
	"log"
	"net"

	"github.com/leegeobuk/go-gRPC/proto/pc"
	"github.com/leegeobuk/go-gRPC/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.String("p", "", "port for server")
	flag.Parse()
	log.Printf("listening to port %s", *port)

	ls := service.NewLaptopServer(service.NewLaptopStore())
	gs := grpc.NewServer()
	pc.RegisterLaptopServiceServer(gs, ls)

	l, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatalf("cannot start the server %v", err)
	}

	err = gs.Serve(l)
	if err != nil {
		log.Fatalf("cannot start the server %v", err)
	}
}
