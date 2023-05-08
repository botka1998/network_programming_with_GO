package main

import (
	ch03 "network_programming_with_GO/ch03/1_TCP"
	"testing"
)

func main() {
	t := &testing.T{}
	ch03.TestListener(t)
	ch03.TestDial(t)
}
