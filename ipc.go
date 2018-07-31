package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	iName := "en0"
	args := os.Args[1:]
	if len(args) > 0 {
		iName = args[0]
	}

	intr, err := net.InterfaceByName(iName)

	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	addrs, err := intr.Addrs()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}

	printed := false
	for _, addr := range addrs {
		if addr, ok := addr.(*net.IPNet); ok {
			if addr.IP.To4() != nil {
				fmt.Println(addr.IP.To4())
				printed = true
			}
		}
	}

	if !printed {
		fmt.Fprintln(os.Stderr, "No ip found on", iName)
		os.Exit(0)
	}
}
