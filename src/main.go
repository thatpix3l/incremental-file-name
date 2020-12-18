package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Flags
	textPtr := flag.String("text", "", "Text to parse.")
	portPtr := flag.String("port", "", "Port number.")
	flag.Parse()

	// Screw around with missing flag values and how to handle them
	if *textPtr == "" || *portPtr == "" {
		fmt.Printf("You didn't provide any options.\nNow you will face the consequences...")
		os.Exit(1)
	}

	fmt.Printf("Text you typed: %s\nPort you typed: %s", *textPtr, *portPtr)
}
