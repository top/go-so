package main

import "C"

type Chinese struct{}

func (Chinese) Say() string {
	return "你好, 世界!"
}

var Lang Chinese
