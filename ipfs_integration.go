package main

import (
	"fmt"
)

type IPFSClient struct {
	Gateway string
}

func NewIPFSClient() *IPFSClient {
	return &IPFSClient{
		Gateway: "https://ipfs.io/ipfs/",
	}
}

func (ipfs *IPFSClient) Upload(data string) string {
	cid := "Qm" + fmt.Sprintf("%x", len(data))
	fmt.Printf("数据已上传 IPFS，CID：%s\n", cid)
	return cid
}

func (ipfs *IPFSClient) Download(cid string) string {
	fmt.Printf("从 IPFS 下载数据：%s\n", cid)
	return "ipfs_data_content"
}
