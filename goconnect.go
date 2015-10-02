package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	protocol := "tcp"
	netaddress := "127.0.0.1:12345"
	if len(os.Args) > 1 {
		netaddress = os.Args[1]
	} else {
		fmt.Println("Note: arguments 'goconnect ip:port protocol' supported")
		fmt.Println("Example: 'goconnect 192.168.1.1:6666 udp'")
		fmt.Println("Using default settings")
	}
	if len(os.Args) > 2 {
		protocol = os.Args[2]
	}

	fmt.Printf("Connecting to %s with protocol %s...", netaddress, protocol)
	conn, err := net.Dial(protocol, netaddress)
	if err != nil {
		fmt.Println("\nError connecting:", err.Error())
		return
	}
	fmt.Printf("Success!\n---DATA---\n")
	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		fmt.Println("Error reading connection:", err.Error())
		return
	}
	fmt.Println(string(reply))
	fmt.Println("\n---DATA---\nSucessfully read data. Exiting.")
}
