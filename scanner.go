package main

import (
	"fmt"
	"os"
	"strings"
)

/*
scanner: Simple scanner to determine whether a search argument is inside a file.
*/
func scanner(fullFilePath string, branchedFile string, searchArg string, caseIgnored bool) int {
	// Get all the bytes of the current selected file
	dataBytes, err := os.ReadFile(fullFilePath)
	if err != nil {
		croak("scanner: os.ReadFile(%s) failed, err:%s\n", fullFilePath, err)
	}

	// Convert bytes into an array of strings
	dataStrings := strings.Split(string(dataBytes), "\n")

	// For each string in the file, see if the search argument is present
	gotAHit := 0
	var argIndex int
	var lcLine, lcArg string
	for lineNumber, line := range dataStrings {
		if caseIgnored {
			lcLine = strings.ToLower(line)
			lcArg = strings.ToLower(searchArg)
			argIndex = strings.Index(lcLine, lcArg)
		} else {
			argIndex = strings.Index(line, searchArg)
		}
		if argIndex > -1 {
			line = strings.ReplaceAll(line, "\n", "")
			fmt.Printf("%s:%d %s\n", branchedFile, lineNumber+1, line)
			gotAHit = 1
		}
	}

	// Return the "got a hit" flag (0 or 1).
	return gotAHit
}
