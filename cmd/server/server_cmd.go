package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"

	"github.com/hashicorp/mdns"
)

func main() {
	host, _ := os.Hostname()
	info := []string{"My awesome service"}
	service, _ := mdns.NewMDNSService(host, "_foobar._tcp", "", "", 8000, nil, info)
	var server *mdns.Server
	/*
		! Q:不绑定网卡的情况下是否能接收mDNS
		! A:windows下必须绑定，linux下可以默认
	*/
	sysType := runtime.GOOS
	switch sysType {
	case "windows":
		ifce, _ := net.InterfaceByName("WLAN")
		server, _ = mdns.NewServer(&mdns.Config{Zone: service, Iface: ifce})
	case "linux":
		server, _ = mdns.NewServer(&mdns.Config{Zone: service})
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	res := <-c
	if res == os.Interrupt {
		fmt.Println("[get signal:]", os.Interrupt)
		server.Shutdown()
	}
}
