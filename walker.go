package main

import (
	"os"
	"path/filepath"
)

func isInList(baseName string, argList []string) bool {
	for _, pattern := range argList {
		match, err := filepath.Match(pattern, baseName)
		if err != nil {
			croak("walker: filepath.Match(%s, %s) failed, err:%s\n", pattern, baseName, err.Error())
		}
		if match {
			return true
		}
	}
	return false
}

/*
walker:
Walk the specified directory tree (pathTreeTop).
Obey the inclusion list (global.incList).
Obey the exclusion list (global.excList).
Scan all qualifying files for the argument (global.arg).
*/
func walker(pathTreeTop string, hitCount, skipCount int) (int, int) {

	// Get pointer to global definitions.
	global := GetGlobalRef()

	// Get all the entries in the current directory level.
	entries, err := os.ReadDir(pathTreeTop)
	if err != nil {
		croak("walker: os.ReadDir(%s) failed, err:%s\n", pathTreeTop, err.Error())
	}

	// For each entry at the current directory level, process it.
	for _, dirEntry := range entries {

		// Get base name for the current entry.
		baseName := dirEntry.Name()

		// Get full path.
		fullPathFile := filepath.Join(pathTreeTop, baseName)

		// Get file info.
		fileInfo, err := os.Stat(fullPathFile)
		if err != nil {
			croak("walker: os.Stat(%s) failed, err:%s\n", fullPathFile, err.Error())
		}

		// If this is a subdirectory entry, it needs special handling.
		if fileInfo.IsDir() {
			// Recur.
			hitCount, skipCount = walker(fullPathFile, hitCount, skipCount)
		}

		// Exclude it?
		if isInList(baseName, global.excList) {
			skipCount += 1
			continue
		}

		// Include it?
		if !isInList(baseName, global.incList) {
			skipCount += 1
			continue
		}

		// This file should be scanned.
		hitCount += scanner(fullPathFile, global.arg)
	}

	// Processed all entries at this level.
	return hitCount, skipCount

}
