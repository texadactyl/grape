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

## Example Execution

Suppose one wishes to find the file and line number of every reference to "sql" in all of the "*.java" files in the tree.
```
grape -i '*.java' -e 'w*.go' -d ~/sour sql

/home/gherkin/sour/essai/db-sqlite/main.java:1 import java.sql.*;
/home/gherkin/sour/essai/db-sqlite/main.java:5     static String connUrl = "jdbc:sqlite:../../database/jacotest.db";   // sqlite only
/home/gherkin/sour/essai/db-sqlite/main.java:6     static String queryString = "SELECT test_case, date_utc, time_utc, result, fail_text FROM history ORDER BY test_case, date_utc DESC, time_utc DESC";   // sqlite only
main: File hit count: 1
main: File skip count: 2654
```
In this example, only one *.java file contained the string "sql" and a total of 2654 files were skipped because they did not have the ".java" file extension.
