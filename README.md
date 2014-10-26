# csvsplit

Split a .csv into multiple files.

## Install

```bash
# The command below requires you to have Go installed
# https://golang.org/doc/install
$ go get github.com/JeffPaine/csvsplit
```

## Examples
```bash
# Basic usage
$ csvsplit -records <number of records> <file>

# Split file.csv into files with 300 records a piece
$ csvplit -records 300 file.csv

# Split file.csv into files with 37 records a piece into the subfolder 'stuff'
$ csvplit -records 37 -output stuff/ file.csv

# Split file.csv into files with 40 records a piece and two header lines
$ csvplit -records 40 -headers 2 file.csv

# Accept csv data from stdin
$ cat file.csv | csvsplit -records 20

# You can use the -output flag to customize the resulting filenames.
# The below will generate custom_filename-001.csv, custom_filename-002.csv, etc.
$ cat file.csv | csvsplit -records 20 -output custom_filename-
```

## Flags
`-records`: Number of records per file  
`-output`: Output filename / path (optional)  
`-headers`: Number of header lines in the input file to add to each ouput file (optional, default=0)
