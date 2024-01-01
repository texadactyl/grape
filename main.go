package main

import (
	"fmt"
	"os"
	"strings"
)

// Show help and then exit to the O/S
func showHelp() {
	_ = InitGlobals()
	fmt.Printf("\nUsage:  grape  [-h]  [-i inclusion-list]  [-e exclusion-list]  arg\n\nwhere\n")
	fmt.Printf("\t-h : This display\n")
	fmt.Printf("\t-i : List of comma-separated globs to include E.g. '*.go,*.java'. Default: '*' (include every file).\n")
	fmt.Printf("\t-e : List of comma-separated globs to exclude E.g. '*_test.go,.*'. Default: '' (no exclusions).\n")
	fmt.Printf("\t-d : Tree-top directory to begin search. Default: current working directory.\n")
	fmt.Printf("\targ : Fixed string argument to search for (required).\n")
	ShowExecInfo()
	os.Exit(0)
}

func main() {
	var Args []string

	// Initialise global.
	global := InitGlobals()

	// Initialise Args to the command-line arguments
	for _, singleVar := range os.Args[1:] {
		Args = append(Args, singleVar)
	}

	// Parse command line arguments
	nargs := len(Args)
	if nargs < 1 {
		fmt.Println("main: No arguments specified.")
		showHelp()
	}
	for ii := 0; ii < nargs; ii++ {
		switch Args[ii] {

		case "-h":
			showHelp()

		case "-i":
			ii += 1
			global.incList = strings.Split(Args[ii], ",")

		case "-e":
			ii += 1
			global.excList = strings.Split(Args[ii], ",")

		case "-d":
			ii += 1
			global.treeTop = Args[ii]
			err := os.Chdir(global.treeTop)
			if err != nil {
				croak("main: os.Chdir(%s) failed, err:%s\n", Args[ii], err.Error())
			}

		default:
			if strings.HasPrefix(Args[ii], "-") {
				croak("main: Unrecognizable option argument: %s", Args[ii])
			}
			global.arg = Args[ii]
			ii += 1
			if ii != nargs {
				croak("main: Something extraneous appears after the search argument: %s", Args[ii])
			}
		}
	}

	hitCount := 0
	skipCount := 0
	hitCount, skipCount = walker(global.treeTop, hitCount, skipCount)
	fmt.Printf("main: File hit count: %d\n", hitCount)
	fmt.Printf("main: File skip count: %d\n", skipCount)
}
