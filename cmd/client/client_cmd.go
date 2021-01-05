package main

import (
	"fmt"

	"github.com/hashicorp/mdns"
)

func main() {
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry: %v\n", entry)
		}
	}()
	// !服务器端暂时对PTR 和 ANY感兴趣，检查当前默认的类型是什么
	// Start the lookup
	mdns.Lookup("_foobar._tcp", entriesCh)
	//! close(entriesCh)
}

// (ip.version==4&&mdns)&&(ip.addr == 224.0.0.251||udp.port==5353||udp.port==8000||ip.addr==172.22.60.132)
