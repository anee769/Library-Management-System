package Library

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func GobEncode(object any) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	gob.Register(PhysicalBook{})
	gob.Register(DigitalBook{})
	if err := encoder.Encode(object); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func GobDecode(data []byte, object any) (any, error) {
	reader := bytes.NewReader(data)
	decoder := gob.NewDecoder(reader)
	if err := decoder.Decode(&object); err != nil {
		return nil, err
	}
	fmt.Println(object)
	return object, nil
}
