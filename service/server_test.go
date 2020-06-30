package service

import (
	"context"
	"testing"

	"github.com/leegeobuk/go-gRPC/proto/pc"
	"github.com/leegeobuk/go-gRPC/sample"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateServer(t *testing.T) {
	t.Parallel()

	laptopWithoutID := sample.NewLaptop()
	laptopWithoutID.ID = ""

	laptopWithInvalidID := sample.NewLaptop()
	laptopWithInvalidID.ID = "invalid-uuid"

	laptopWithDuplicateID := sample.NewLaptop()
	storeWithDuplicateID := NewLaptopStore()
	err := storeWithDuplicateID.Save(laptopWithDuplicateID)
	require.Nil(t, err)

	testCases := []struct {
		name   string
		laptop *pc.Laptop
		db     DB
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			db:     NewLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopWithoutID,
			db:     NewLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invalid_id",
			laptop: laptopWithInvalidID,
			db:     NewLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_duplicate_id",
			laptop: laptopWithDuplicateID,
			db:     storeWithDuplicateID,
			code:   codes.AlreadyExists,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pc.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			server := NewLaptopServer(tc.db)
			res, err := server.CreateLaptop(context.Background(), req)

			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.ID)
				if len(tc.laptop.ID) > 0 {
					require.Equal(t, tc.laptop.ID, res.ID)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)

				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, s.Code())
			}
		})
	}
}
