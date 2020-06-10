package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// ToBinary writes protocol buffer message to binary file
func ToBinary(m proto.Message, filename string) error {
	data, err := proto.Marshal(m)

	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)

	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}

	return nil
}

// FromBinary reads protocol buffer message from binary file
func FromBinary(m proto.Message, filename string) error {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return fmt.Errorf("cannot read binary file : %w", err)
	}

	err = proto.Unmarshal(data, m)

	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %w", err)
	}

	return nil
}
