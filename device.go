package main

type Sender interface {
	Send([]byte)
}

type Receiver interface {
	Receive([]byte)
}

type Shell interface {
	Run()
}

// Device combines the previous interfaces
type Device interface {
	Sender
	Receiver
	Shell
}

