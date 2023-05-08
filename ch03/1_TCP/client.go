package ch03

import (
	"net"
	"testing"
)

func TestDial(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	conn.Close()
}
