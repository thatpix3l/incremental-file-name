package main

import (
	"fmt"
	"os"
	"flag"
	"path/filepath"
)

func main() {
	fileNamePtr := flag.String("file", "", "Path to file")
	flag.Parse()
	fileName := *fileNamePtr

	i := 0
	for fileExists(fileName) {
		fmt.Println(filepath.Base(fileName))
		fileName = "literally anything else"
		i++
	}
}

func fileExists(fileInput string) bool {
	_, errStat := os.Stat(fileInput)
  boolStatus := true
	if errStat != nil {
		boolStatus = false
	}
	return boolStatus
}