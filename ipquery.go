package gobelt

import (
	"fmt"
	"net"
)

func IPQuery() {
	addrs, _ := net.InterfaceAddrs()
	fmt.Printf("%v\n", addrs)
	for _, addr := range addrs {
		fmt.Println("IPv4: ", addr)
	}
}
