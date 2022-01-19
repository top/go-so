package main

import (
	"fmt"
	"os"
	"plugin"
)

type Language interface {
	Say() string
}

func pluginCall(fn string) {
	p, err := plugin.Open(fn)
	if err != nil {
		fmt.Println("plugin open failed:", fn, err)
		return
	}

	l, err := p.Lookup("Lang")
	if err != nil {
		fmt.Println("plugin lookup failed:", err)
		return
	}

	h, ok := l.(Language)
	if !ok {
		fmt.Println("plugin execution failed")
		return
	}

	fmt.Println(fn, h.Say())
}

func main() {
	if len(os.Args) == 2 {
		pluginCall(os.Args[1])
	}
}
