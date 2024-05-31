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
	rdphostquery := flag.Bool("rdphostquery", false, "Queries registry for successful RDP sessions performed by Current User")
	version := flag.Bool("version", false, "Prints version information")
	flag.Parse()

	if *rdphostquery {
		key, e = gobelt.RDPHostQuery()
		if e != nil {
			fmt.Printf("Received an error %s", e.Error())
		} else {
			fmt.Printf("[+] The value of the registry key is %s", key)
		}
	}
	if *version {
		fmt.Printf("[+] Gobelt version 1.0\n")
	}

}
