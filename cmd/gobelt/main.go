package main

import (
	"flag"
	"fmt"
	"gobelt"
)

func main() {
	var key []string
	var e error
	fmt.Print("[+] Welcome to Gobelt, the Golang alternative to Seatbelt\n")
	// fun with flags
	rdpregquery := flag.Bool("rdpregquery", false, "Queries registry for successful RDP sessions performed by Current User")
	ipquery := flag.Bool("ipquery", false, "Queries ip addressing information of network interfaces")
	version := flag.Bool("version", false, "Prints version information")
	flag.Parse()

	if *rdpregquery {
		key, e = gobelt.RDPRegQuery()
		if e != nil {
			fmt.Printf("Received an error %s", e.Error())
		} else {
			fmt.Printf("[+] The value of the registry key is %s", key)
		}
	}
	if *ipquery {
		fmt.Printf("[+] Printing out list of IPv4 IPs that are currently assigned to an interface\n")
		addr := gobelt.IPQuery()
		for i := 0; i < len(addr); i++ {
			fmt.Println(addr[i])
		}
	}
	if *version {
		fmt.Printf("[+] Gobelt version 1.0\n")
	}

}
