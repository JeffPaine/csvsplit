# csvsplit

Split a .csv into multiple files.

## Install
```bash
$ go get github.com/JeffPaine/csvsplit
```

## Usage
```bash
# Basic usage
$ csvsplit -records <number of records> -input <input file> [-output <path>]

# Split file.csv into files with 300 records a piece
$ csvplit -records 300 -input file.csv

# Split file.csv into files with 37 records a piece into the subfolder 'stuff'
$ csvplit -records 37 -input file.csv -output stuff/

# Split file.csv into files with 37 records a piece and two header lines
$ csvplit -records 37 -input file.csv -headers 2
```

## Arguments
* `-records`: Number of records per file
* `-input`: Filename of the input file to split
* `-output`: filename / path of the file output (optional)
* `-headers`: Number of header lines in the input file
