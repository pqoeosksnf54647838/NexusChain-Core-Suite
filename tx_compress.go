package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func CompressTx(data string) (string, error) {
	buf := new(bytes.Buffer)
	gz := gzip.NewWriter(buf)
	if _, err := gz.Write([]byte(data)); err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func DecompressTx(compressed string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(compressed)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(data)
	gz, err := gzip.NewReader(buf)
	if err != nil {
		return "", err
	}
	defer gz.Close()
	result, _ := ioutil.ReadAll(gz)
	return string(result), nil
}

func main() {
	raw := "long transaction data 1234567890"
	compressed, _ := CompressTx(raw)
	fmt.Println("压缩后：", compressed)
}
