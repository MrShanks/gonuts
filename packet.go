package main

import "net"

type Packet struct {
	SrcIP   net.IP
	DstIP   net.IP
	Payload []byte
}

func NewPacket(srcIP, dstIP net.IP, payload []byte) *Packet {
	return &Packet{
		SrcIP:   srcIP,
		DstIP:   dstIP,
		Payload: payload,
	}
}
