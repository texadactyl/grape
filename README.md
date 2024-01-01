# grape
Abbreviated Gnu grep for Windows

## Installation

```
git clone https://github.com/texadactyl/grape
cd grape
go install (preferred) or go build into a directory in the PATH
```

## Help

```
Usage:  grape  [-h]  [-i inclusion-list]  [-e exclusion-list]  arg

where
	-h : This display
	-i : List of comma-separated globs to include E.g. '*.go,*.java'. Default: '*' (include every file).
	-e : List of comma-separated globs to exclude E.g. '*_test.go,.*'. Default: '' (no exclusions).
	-d : Tree-top directory to begin search. Default: current working directory.
	arg : Fixed string argument to search for (required).
grape version: x.y.z
Built with: gox.y.z
```
