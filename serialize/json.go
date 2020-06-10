package serialize

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// ToJSON writes protocol buffer message to binary file
func ToJSON(m proto.Message, filename string) error {
	data, err := convertToJSON(m)

	if err != nil {
		return fmt.Errorf("cannot marshal proto message jo JSON: %w", err)
	}

	err = ioutil.WriteFile(filename, []byte(data), 0644)

	if err != nil {
		return fmt.Errorf("cannot write data to binary file: %w", err)
	}

	return nil
}

// convertToJSON converts protocol buffer message to JSON string
func convertToJSON(m proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       " ",
		OrigName:     true,
	}
	return marshaler.MarshalToString(m)
}
