package encoding

import (
	"encoding/json"
	"log"
)

// Currently we are using JSON
// If we have to use XML or any other format
// then we have to just make changes in this encoding package

type Input struct {
	Key string `json:"key"`
}

func (input *Input) ReadFromBytes(msg []byte) error {
	err := json.Unmarshal(msg, input)
	if err != nil {
		log.Println("Error during Unmarshal", err)
		return err
	}
	return nil
}

func (input *Input) WriteToBytes() ([]byte, error) {
	outputMarshal, err := json.Marshal(*input)
	if err != nil {
		log.Println("Error during Marshal", err)
		return outputMarshal, err
	}

	return outputMarshal, nil
}

type Output struct {
	Value string `json:"value"`
}

func (output *Output) ReadFromBytes(msg []byte) error {
	err := json.Unmarshal(msg, output)
	if err != nil {
		log.Println("Error during Unmarshal", err)
		return err
	}
	return nil
}

func (output *Output) WriteToBytes() ([]byte, error) {
	outputMarshal, err := json.Marshal(*output)
	if err != nil {
		log.Println("Error during Marshal", err)
		return outputMarshal, err
	}

	return outputMarshal, nil
}
