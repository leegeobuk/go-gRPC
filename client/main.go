package main

import (
	"context"
	"flag"
	"log"

	"github.com/leegeobuk/go-gRPC/proto/pc"
	"github.com/leegeobuk/go-gRPC/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	port := flag.String("p", "", "server port")
	flag.Parse()
	log.Printf("dial server port %s", *port)

	// connecting to server
	conn, err := grpc.Dial("localhost:"+*port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server %v", err)
	}

	lc := pc.NewLaptopServiceClient(conn)

	// createing request
	laptop := sample.NewLaptop()
	req := &pc.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := lc.CreateLaptop(context.Background(), req)
	if err != nil {
		status, ok := status.FromError(err)
		if ok && status.Code() == codes.AlreadyExists {
			log.Print("laptop already exists!")
		} else {
			log.Fatalf("cannot create laptop %v", err)
		}
		return
	}

	log.Printf("laptop created with id: %s", res.ID)
}
