package main

import (
	"fmt"
	"gobelt"
)

func main() {
	var key []string
	var e error
	fmt.Print("[+] Welcome to gobelt, the Golang alternative to Seatbelt\n")
	key, e = gobelt.RDPHostQuery()
	if e != nil {
		fmt.Printf("Received an error %s", e.Error())
	} else {
		fmt.Printf("[+] The value of the registry key is %s", key)
	}
}
