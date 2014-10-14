// csvsplit: Split a .csv into multiple files.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var records = flag.Int("records", 0, "The number of records per file")
var input = flag.String("input", "", "Filename of the input file to split")
var output = flag.String("output", "", "filename / path of the file output (optional)")
var headerLines = flag.Int("headers", 1, "Number of header lines in the input file (will be repeated in each output file")

func main() {
	flag.Parse()
	if *input == "" || *records < 1 || *headerLines < 0 {
		flag.Usage()
		os.Exit(1)
	}

	csvFile, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	recordsToWrite := make([][]string, 0)
	fileCount := 1
	headers := make([][]string, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if len(headers) < *headerLines {
			headers = append(headers, record)
			continue
		}

		recordsToWrite = append(recordsToWrite, record)
		if len(recordsToWrite) == *records-*headerLines {
			saveCSVFile(headers, recordsToWrite, fileCount)
			recordsToWrite = make([][]string, 0)
			fileCount += 1
		}
	}
	if len(recordsToWrite) > 0 {
		saveCSVFile(headers, recordsToWrite, fileCount)
	}
}

func saveCSVFile(h [][]string, r [][]string, fileCount int) {
	fileName := fmt.Sprintf("%v%03d%v", *output, fileCount, ".csv")
	if _, err := os.Stat(fileName); err == nil {
		log.Fatal("File exists: ", fileName)
	}
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	writer := csv.NewWriter(f)
	if len(h) > 0 {
		writer.WriteAll(h)
	}
	writer.WriteAll(r)
}
