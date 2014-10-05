# csvsplit

Split a .csv file into multiple files with a given number of records per file.

## Usage
```
csvsplit -r records -i input [-o path]
```

## Install
```bash
$ go get github.com/JeffPaine/csvsplit
```

## Arguments
* `-r`: Number of records per file
* `-i`: Filename of the input file to split
* `-o`: filename / path of the file output (optional)

## Examples
```bash
# Split file.csv into files with 300 records a piece
$ csvplit -r 300 -i file.csv
```
