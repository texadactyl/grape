package main

import (
	"fmt"
	"os"
	"strings"
)

/*
scanner: Simple scanner to determine whether a search argument is inside a file.
*/
func scanner(filePath string, searchArg string) int {
	// Get all the bytes of the current selected file
	dataBytes, err := os.ReadFile(filePath)
	if err != nil {
		croak("scanner: os.ReadFile(%s) failed, err:%s\n", filePath, err)
	}

	// Convert bytes into an array of strings
	dataStrings := strings.Split(string(dataBytes), "\n")

	// For each string in the file, see if the search argument is present
	gotAHit := 0
	for ix, line := range dataStrings {
		if strings.Index(line, searchArg) > -1 {
			line = strings.ReplaceAll(line, "\n", "")
			fmt.Printf("%s:%d %s\n", filePath, ix+1, line)
			gotAHit = 1
		}
	}

	// Return the "got a hit" flag (0 or 1).
	return gotAHit
}
