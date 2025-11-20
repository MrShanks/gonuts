package main

import "net"

type Nic struct {
	Name string
	IP   net.IP
	MAC  net.HardwareAddr
}
