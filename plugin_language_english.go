package main

import "C"

type English struct{}

func (English) Say() string {
	return "Hello, world!"
}

var Lang English
