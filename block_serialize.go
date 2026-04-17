package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func SerializeBlock(block Block) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func DeserializeBlock(data []byte) (Block, error) {
	var block Block
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(&block)
	if err != nil {
		return Block{}, err
	}
	return block, nil
}

func main() {
	block := Block{Index: 1, Data: "test"}
	data, _ := SerializeBlock(block)
	fmt.Printf("序列化长度：%d 字节\n", len(data))
	decoded, _ := DeserializeBlock(data)
	fmt.Printf("反序列化结果：%+v\n", decoded)
}
