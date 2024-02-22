package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go func() {
			for {
				if _, err := io.WriteString(conn, time.Now().Format("15:05:00\n")); err != nil {
					fmt.Println(err)
				}

				time.Sleep(time.Second)
			}
		}()
	}
}
