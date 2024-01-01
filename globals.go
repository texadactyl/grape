package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

const VERSION = "1.0.0"

// Definition of the singleton global.
type GlobalsStruct struct {
	version string   // Software version string
	incList []string // Inclusion list of globs
	excList []string // Exclusion list of globs
	treeTop string   // Tree top directory to begin walk
	arg     string   // Search argument
}

// Here's the singleton.
var global GlobalsStruct

// Convert a UTC time string into a local one
func UTCTimeStr2LocalTimeStr(utcString string) string {
	timeStamp, err := time.Parse("2006-01-02T15:04:05Z07:00", utcString)
	if err != nil {
		croak("UTCTimeStr2LocalTimeStr: time.Parse(2006-01-02T15:04:05Z07:00, %s) error: %s\n", utcString, err.Error())
	}
	zone, _ := time.Now().Zone()
	return fmt.Sprintf("%s %s", timeStamp.Local().Format("2006-01-02 15:04:05"), zone)
}

// Show executable binary information relevant to "vcs"
func ShowExecInfo() {
	fmt.Printf("grape version: %s\n", global.version)
	fmt.Printf("Built with: %s\n", runtime.Version())
	// Only interested in the "vcs." information
	info, _ := debug.ReadBuildInfo()
	for ii := 0; ii < len(info.Settings); ii++ {
		biKey := info.Settings[ii].Key
		biValue := info.Settings[ii].Value
		if strings.HasPrefix(biKey, "vcs.") {
			if biKey == "vcs.time" {
				biValue = UTCTimeStr2LocalTimeStr(biValue)
			}
			fmt.Printf("BuildData %s: %s\n", biKey, biValue)
		}
	}
}

// Initialise the singleton global
func InitGlobals() *GlobalsStruct {
	cwd, err := os.Getwd()
	if err != nil {
		croak("InitGlobals: os.Getwd() failed, err:%s\n", err.Error())
	}
	global = GlobalsStruct{
		version: VERSION,
		incList: []string{"*"},
		excList: []string{},
		treeTop: cwd,
	}
	return &global
}

// GetGlobalRef returns a pointer to the singleton instance of GlobalsStruct
func GetGlobalRef() *GlobalsStruct {
	return &global
}

// Display message and die.
func croak(format string, args ...any) {
	fmt.Printf(format, args)
	os.Exit(1)
}
