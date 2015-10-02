package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	protocol := "tcp"
	netaddress := "127.0.0.1:12345"
	if len(os.Args) > 1 {
		netaddress = os.Args[1]
	} else {
		fmt.Println("Note: arguments 'golisten ip:port protocol' supported")
		fmt.Println("Example: 'golisten 192.168.1.1:6666 udp'")
		fmt.Println("Using default settings")
	}
	if len(os.Args) > 2 {
		protocol = os.Args[2]
	}

	fmt.Println("Listening on", netaddress, "with protocol", protocol)
	psock, err := net.Listen(protocol, netaddress)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	for {
		conn, err := psock.Accept()
		if err != nil {
			fmt.Println("Error on accept:", err.Error())
			return
		}
		fmt.Println(conn.RemoteAddr().String(), "just connected, sending success and close")
		io.Copy(conn, bytes.NewBufferString("Connection successful to golisten!"))
		conn.Close()
	}

}
