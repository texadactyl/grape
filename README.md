# grape
Abbreviated Gnu grep, primarily for Windows.

I created this repository to assist Windows developers who do not wish to install the Gnu toolset on Windows.

## Installation

```
git clone https://github.com/texadactyl/grape
cd grape
go install (preferred) or go build into a directory in the PATH
```

## Help

```
Usage:  grape  [-h]  [-i]  [-n inclusion-list]  [-x exclusion-list]  [-d directory]  arg

where
	-h : This display
	-i : The case of alphabetics is ignored (Abc=ABC). Default: Abc != ABC.
	-n : List of comma-separated globs to include E.g. '*.go,*.java'. Default: '*' (include every file).
	-x : List of comma-separated globs to exclude E.g. '*_test.go,.*'. Default: '' (no exclusions).
	-d : Tree-top directory to begin search. Default: current working directory.
	arg : Fixed string argument to search for (required).
grape version: x.y.z
Built with: gox.y.z
```

## Example Execution

Suppose one wishes to find the file and line number of every reference to "sql" in all of the "*.java" files in the tree.
```
grape -n '*.java' -x 'w*.go' -d ~/sour sql

essai/db-sqlite/main.java:1 import java.sql.*;
essai/db-sqlite/main.java:5     static String connUrl = "jdbc:sqlite:../../database/jacotest.db";   // sqlite only
sour/essai/db-sqlite/main.java:6     static String queryString = "SELECT test_case, date_utc, time_utc, result, fail_text FROM history ORDER BY test_case, date_utc DESC, time_utc DESC";   // sqlite only

Total file hit count: 1
Total file skip count: 2654
```
In this example, only one *.java file contained the string "sql" and a total of 2654 files were skipped because they did not have the ".java" file extension.
