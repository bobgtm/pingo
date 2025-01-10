package main

import (
	"fmt"
	"log"
	"net"
)

// take a url or ip address
// establish conn from host to remote url/ip
// Discover the remote hostname and associated ip
// send stream of raw/empty data
// Check for a return
// Print response from the remote host

func getLocalAdress() net.IP {
	con, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		log.Fatal(error)
		defer con.Close()
	}
	localAddress := con.LocalAddr().(*net.UDPAddr)

	return localAddress.IP

}

func main() {
	ipString := getLocalAdress()
	fmt.Println(ipString)
}
