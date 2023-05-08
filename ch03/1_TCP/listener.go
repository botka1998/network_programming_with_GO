package ch03

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = listener.Close()
	}()
	print("bound to ", listener.Addr().String())
	t.Logf("bound to %q", listener.Addr())

	done := make(chan struct{})
	go func() {
		defer func() { done <- struct{}{} }()

		for {
			c, err := listener.Accept()
			if err != nil {
				t.Log(err)
				return
			}
			go func(c net.Conn) {
				defer func() {
					c.Close()
					done <- struct{}{}
				}()

				buf := make([]byte, 1024)
				for {
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}
					t.Logf("recieved: %q", buf[:n])
					fmt.Printf("recieved: %q", buf[:n])
				}
			}(c)
		}

	}()
	conn, err := net.Dial("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	conn.Close()
	<-done
	listener.Close()
	<-done
}
