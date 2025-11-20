package main

import "net"

type Route struct {
	Destination *net.IPNet
	Gateway     net.IP
	NicName     string
	Metric      int
}

func NewRoute(network net.IP, mask net.IPMask, gateway net.IP, nicName string, priority int) *Route {
	return &Route{
		Destination: &net.IPNet{
			IP:   network,
			Mask: mask,
		},
		Gateway: gateway,
		NicName: nicName,
		Metric:  priority,
	}
}
