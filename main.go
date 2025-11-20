package main

import "fmt"

func main() {
	h := NewHost("host-A")
	h.AddNic("eth0")
	fmt.Println(h)

	h.Run()
}
