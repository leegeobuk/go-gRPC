package serialize

import (
	"testing"

	"github.com/leegeobuk/go-gRPC/pb/pc"
	"github.com/leegeobuk/go-gRPC/sample"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestBinary(t *testing.T) {
	t.Parallel()

	binaryFile := "../temp/laptop.bin"
	jsonFile := "../temp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := ToBinary(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pc.Laptop{}
	err = FromBinary(laptop2, binaryFile)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = ToJSON(laptop1, jsonFile)
	require.NoError(t, err)
}
