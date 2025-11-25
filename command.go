package main

type CommandFunc func(h *Host, args []string)

type Command struct {
	Name        string
	Description string
	Run         CommandFunc
}
