# csvsplit

Split a .csv file into multiple files.

## Install
```bash
$ go get github.com/JeffPaine/csvsplit
```

## Usage
```bash
# Basic usage
$ csvsplit -r <number of records> -i <input file> [-o <path>]

# Split file.csv into files with 300 records a piece
$ csvplit -r 300 -i file.csv

# Split file.csv into files with 37 records a piece into the subfolder 'stuff'
$ csvplit -r 37 -i file.csv -o stuff/
```

## Arguments
* `-r`: Number of records per file
* `-i`: Filename of the input file to split
* `-o`: filename / path of the file output (optional)
