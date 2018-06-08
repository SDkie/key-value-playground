package encoding

import (
	"encoding/json"
	"log"
)

type Input struct {
	Key string `json:"key"`
}

func (input *Input) ReadFromByte(msg []byte) error {
	err := json.Unmarshal(msg, input)
	if err != nil {
		log.Println("Error during Unmarshal", err)
		return err
	}
	return nil
}

func (input *Input) WriteToByte() ([]byte, error) {
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

func (output *Output) ReadFromByte(msg []byte) error {
	err := json.Unmarshal(msg, output)
	if err != nil {
		log.Println("Error during Unmarshal", err)
		return err
	}
	return nil
}

func (output *Output) WriteToByte() ([]byte, error) {
	outputMarshal, err := json.Marshal(*output)
	if err != nil {
		log.Println("Error during Marshal", err)
		return outputMarshal, err
	}

	return outputMarshal, nil
}
