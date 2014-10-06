/*
csvsplit: Split a .csv file into multiple files with a given number of records per file.

Usage
    # Basic usage
    $ csvsplit -r <number of records> -i <input file> [-o path]

    # Split file.csv into files with 300 records a piece
    $ csvplit -r 300 -i file.csv

    # Split file.csv into files with 37 records a piece into the subfolder 'stuff'
    $ csvplit -r 37 -i file.csv -o stuff/

Arguments
    -r: Number of records per file
    -i: Filename of the input file to split
    -o: filename / path of the file output (optional)

*/
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var records = flag.Int("r", 1000, "The number of records per file")
var input = flag.String("i", "", "Filename of the input file to split")
var output = flag.String("o", "", "filename / path of the file output (optional)")

func main() {
	flag.Parse()
	if *input == "" || *records < 1 {
		flag.Usage()
		log.Fatal()
	}

	csvFile, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	recordsToWrite := make([][]string, 0)
	recordCount := 0
	fileCount := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		recordsToWrite = append(recordsToWrite, record)
		recordCount += 1
		if recordCount == *records {
			saveCSVFile(recordsToWrite, fileCount)
			recordCount = 0
			recordsToWrite = make([][]string, 0)
			fileCount += 1
		}
	}
	if len(recordsToWrite) > 0 {
		saveCSVFile(recordsToWrite, fileCount)
	}
}

func saveCSVFile(r [][]string, fileCount int) {
	fileName := fmt.Sprintf("%v%03d%v", *output, fileCount, ".csv")
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		writer := csv.NewWriter(f)
		writer.WriteAll(r)
	} else {
		log.Fatal("File exists: ", fileName)
	}
}
