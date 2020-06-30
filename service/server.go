package service

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/leegeobuk/go-gRPC/proto/pc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopServer provides laptop services
type LaptopServer struct {
	DB DB
}

// NewLaptopServer returns a new LaptopServer
func NewLaptopServer(db DB) *LaptopServer {
	return &LaptopServer{db}
}

// CreateLaptop is a unary RPC to create a new laptop
func (s *LaptopServer) CreateLaptop(ctx context.Context, req *pc.CreateLaptopRequest) (
	*pc.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("CreateLaptopRequest received with id: %s", laptop.ID)

	if len(laptop.ID) > 0 {
		_, err := uuid.Parse(laptop.ID)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop id is not a valid uuid: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()

		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate new laptop id: %v", err)
		}

		laptop.ID = id.String()
	}

	err := s.DB.Save(laptop)
	if err != nil {
		code := codes.Internal

		if errors.Is(err, ErrorExistingID) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop to db: %v", err)
	}

	log.Printf("laptop saved with id: %s", laptop.ID)
	res := &pc.CreateLaptopResponse{ID: laptop.ID}
	return res, nil
}
