package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type Host struct {
	Name         string
	Nics         map[string]Nic
	RoutingTable []*Route
	Commands     map[string]Command
}

func NewHost(name string) *Host {
	defaultRoute := NewRoute(net.ParseIP("0.0.0.0"), net.IPv4Mask(0, 0, 0, 0), net.ParseIP("192.168.1.1"), "eth0", 1)
	routingTable := make([]*Route, 0)
	routingTable = append(routingTable, defaultRoute)

	h := &Host{
		Name:         name,
		Nics:         make(map[string]Nic),
		RoutingTable: routingTable,
		Commands:     make(map[string]Command),
	}

	h.RegisterCommand("route", "Show routing table", printRoute)
	h.RegisterCommand("hostname", "Show host name", printHostname)
	h.RegisterCommand("ip", "Manage network interfaces", printAddresses)

	return h
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

		input := scanner.Text()
		if strings.TrimSpace(input) == "" {
			continue
		}

		parts := strings.Fields(input)
		cmdName := parts[0]
		args := parts[1:]

		if cmdName == "exit" {
			os.Exit(0)
		}

		if cmd, exists := h.Commands[cmdName]; exists {
			cmd.Run(h, args)
		} else {
			printHelp(h, args)
		}
	}
}

func (h *Host) RegisterCommand(name, desc string, cmd CommandFunc) {
	h.Commands[name] = Command{
		Name:        name,
		Description: desc,
		Run:         cmd,
	}
}

func (h *Host) AddNic(name string) {
	mac, err := generateRandomMAC()
	if err != nil {
		fmt.Printf("Error generating new Hardware address: %v\n", err)
		return
	}

	h.Nics[name] = Nic{
		MAC: mac,
	}
}

func (h *Host) AddRoute(addr net.IP, mask net.IPMask, gw net.IP, nic string, priority int) {
	if _, ok := h.Nics[nic]; !ok {
		fmt.Printf("interface %s, does not exist\n", nic)
		return
	}
	route := NewRoute(addr, mask, gw, nic, priority)
	h.RoutingTable = append(h.RoutingTable, route)
}

func generateRandomMAC() (net.HardwareAddr, error) {
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

func printHelp(h *Host, args []string) {
	fmt.Printf("\nCommand not found\n")
	fmt.Printf("Refer to the list below for the available commands:\n\n")
	for _, cmd := range h.Commands {
		fmt.Printf("- %-20s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
}

func printRoute(h *Host, args []string) {
	for _, r := range h.RoutingTable {
		fmt.Printf("%v\n", r)
	}
}

func printHostname(h *Host, args []string) {
	fmt.Println(h.Name)
}

func printAddresses(h *Host, args []string) {
	if len(args) == 0 {
		return
	}
	if args[0] == "addr" {
		for name, nic := range h.Nics {
			fmt.Printf("interface: %s, mac: %s\n", name, nic.MAC.String())
		}
	}
	if args[0] == "r" {
		findRoute()
	}
}

func findRoute() {
	fmt.Printf("Finding the best route!\n")
}
