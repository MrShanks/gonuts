package main

import (
	"fmt"
	"net"
)

func main() {
	h := NewHost("host-A")
	h.AddNic("eth0")
	h.AddNic("eth1")
	h.AddNic("eth99")

	_, network, err := net.ParseCIDR("10.10.10.0/24")
	if err != nil {
		fmt.Println(err)
	}

	h.AddRoute(network.IP, network.Mask, net.ParseIP("10.10.10.1"), "eth1", 10)

	h.Run()
}
