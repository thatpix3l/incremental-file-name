package main

import (
	"flag"
	"fmt"
)

func main() {
	textPtr := flag.String("text", "", "Text to parse.")
	flag.Parse()

	fmt.Printf("%s", *textPtr)
}
