package main

import (
	"fmt"
	"os"
	"flag"
	"path/filepath"
	"strings"
	"strconv"
)

func main() {
	oldFileNamePtr := flag.String("file", "", "Path to file")
	flag.Parse()

	fileNameBase := strings.TrimSuffix(*oldFileNamePtr, filepath.Ext(*oldFileNamePtr))
	fileNameExt := filepath.Ext(*oldFileNamePtr)
	newFileName := *oldFileNamePtr

	for i := 0; fileExists(newFileName); i++ {
		newFileName = fileNameBase+" "+strconv.Itoa(i)+fileNameExt
	}
	fmt.Println(newFileName)
}

func fileExists(fileInput string) bool {
	_, errStat := os.Stat(fileInput)
  boolStatus := true
	if errStat != nil {
		boolStatus = false
	}
	return boolStatus
}