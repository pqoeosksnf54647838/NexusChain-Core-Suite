package main

import (
	"fmt"
	"net"
)

type DNSSeed struct {
	SeedDomains []string
	NodeList    []string
}

func NewDNSSeed() *DNSSeed {
	return &DNSSeed{
		SeedDomains: []string{"seed1.chain", "seed2.chain"},
		NodeList:    []string{},
	}
}

func (d *DNSSeed) Resolve() {
	fmt.Println("DNS 种子节点解析中...")
	for _, domain := range d.SeedDomains {
		ips, _ := net.LookupIP(domain)
		for _, ip := range ips {
			node := ip.String() + ":9000"
			d.NodeList = append(d.NodeList, node)
			fmt.Println("解析节点：", node)
		}
	}
}
