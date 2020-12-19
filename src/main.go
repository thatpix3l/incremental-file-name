package main

import (
	"fmt"
	"os"
	"github.com/akamensky/argparse"
	"path/filepath"
	"strings"
	"strconv"
)

func main() {
	parser := argparse.NewParser("incremental file name", "Generates a new file name from user input")
	oldFileNamePtr := parser.String("f", "file", &argparse.Options{Required: true, Help: "File name to modify"})
	parserErr := parser.Parse(os.Args)

	if parserErr != nil {
		fmt.Println(parser.Usage(parserErr))
	}

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