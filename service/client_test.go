package service

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/leegeobuk/go-gRPC/proto/pc"
	"github.com/leegeobuk/go-gRPC/sample"
	"github.com/leegeobuk/go-gRPC/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	ls, addr := startLaptopServer(t)
	lc := newLaptopClient(t, addr)

	laptop := sample.NewLaptop()
	expectedID := laptop.ID

	req := &pc.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := lc.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.ID)

	// find and check if the laptop is nil
	other, err := ls.DB.Find(res.ID)
	require.NoError(t, err)
	require.NotNil(t, other)

	isEqual, _ := compareLaptop(laptop, other)
	if !isEqual {
		t.Error("returned laptop is not equal to given laptop %w")
	}
}

func startLaptopServer(t *testing.T) (ls *LaptopServer, addr string) {
	ls = NewLaptopServer(NewLaptopStore())

	gs := grpc.NewServer()
	pc.RegisterLaptopServiceServer(gs, ls)

	l, err := net.Listen("tcp", ":8080")
	require.NoError(t, err)
	addr = l.Addr().String()

	go gs.Serve(l)

	return
}

func newLaptopClient(t *testing.T, addr string) pc.LaptopServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(t, err)
	return pc.NewLaptopServiceClient(conn)
}

func compareLaptop(laptop1, laptop2 *pc.Laptop) (bool, error) {
	json1, err := serializer.ConvertToJSON(laptop1)
	if err != nil {
		return false, fmt.Errorf("cannot convert laptop to json %w", err)
	}

	json2, err := serializer.ConvertToJSON(laptop1)
	if err != nil {
		return false, fmt.Errorf("cannot convert laptop to json %w", err)
	}

	return json1 == json2, nil
}
