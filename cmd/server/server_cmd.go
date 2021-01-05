package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/hashicorp/mdns"
)

func main() {
	host, _ := os.Hostname()
	info := []string{"My awesome service"}
	service, _ := mdns.NewMDNSService(host, "_foobar._tcp", "", "", 8000, nil, info)

	ifce, _ := net.InterfaceByName("WLAN")
	server, _ := mdns.NewServer(&mdns.Config{Zone: service, Iface: ifce})
	// server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	res := <-c
	if res == os.Interrupt {
		fmt.Println("[get signal:]", os.Interrupt)
		server.Shutdown()
	}
}
