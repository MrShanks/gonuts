package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"net"
	"os"
)

type Host struct {
	Name         string
	Nics         map[string]Nic
	RoutingTable []Route
}

func (h *Host) Send([]byte) {}

func (h *Host) Receive([]byte) {}

func (h *Host) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s$ ", h.Name)
		if !scanner.Scan() {
			break // Handle EOF (Ctrl+D)
		}

		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		switch input := scanner.Text(); input {
		case "route":
			for _, r := range h.RoutingTable {
				fmt.Printf("%v\n", r)
			}
		case "hostname":
			fmt.Println(h.Name)
		case "ip addr":
			for name, nic := range h.Nics {
				fmt.Printf("interface: %s, mac: %s\n", name, nic.MAC.String())
			}
		case "exit":
			os.Exit(0)
		}
	}
}

func NewHost(name string) *Host {
	defaultRoute := NewRoute(net.ParseIP("0.0.0.0"), net.IPv4Mask(0, 0, 0, 0), net.ParseIP("192.168.1.1"), "eth0", 1)
	routingTable := []Route{*defaultRoute}

	return &Host{
		Name:         name,
		Nics:         make(map[string]Nic),
		RoutingTable: routingTable,
	}
}

func (h *Host) AddNic(name string) {
	mac, err := GenerateRandomMAC()
	if err != nil {
		fmt.Printf("Error generating new Hardware address: %v\n", err)
		return
	}

	h.Nics[name] = Nic{
		MAC: mac,
	}
}

func GenerateRandomMAC() (net.HardwareAddr, error) {
	buf := make([]byte, 6)

	// Fill the buffer with cryptographically secure random bytes
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}

	// Modify the first byte to respect MAC standards:
	// Set the "Locally Administered" bit (2nd least significant bit) to 1.
	buf[0] |= 0x02

	// Ensure the "Multicast" bit (least significant bit) is 0.
	buf[0] &= 0xfe

	return net.HardwareAddr(buf), nil
}
