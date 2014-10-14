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

var (
	flagRecords = flag.Int("records", 0, "The number of records per file")
	flagInput = flag.String("input", "", "Filename of the input file to split")
	flagOutput = flag.String("output", "", "filename / path of the file output (optional)")
	flagHeaders = flag.Int("headers", 1, "Number of header lines in the input file (will be repeated in each output file")
)

func main() {
	flag.Parse()
	if *flagInput == "" || *flagRecords < 1 || *flagHeaders < 0 {
		flag.Usage()
		os.Exit(1)
	}

	csvFile, err := os.Open(*flagInput)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	headers := make([][]string, 0)
	records := make([][]string, 0)
	fileCount := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if len(headers) < *flagHeaders {
			headers = append(headers, record)
			continue
		}

		records = append(records, record)
		if len(records) == *flagRecords-*flagHeaders {
			saveCSVFile(headers, records, fileCount)
			records = make([][]string, 0)
			fileCount += 1
		}
	}
	if len(records) > 0 {
		saveCSVFile(headers, records, fileCount)
	}
}

func saveCSVFile(h [][]string, r [][]string, fileCount int) {
	fileName := fmt.Sprintf("%v%03d%v", *flagOutput, fileCount, ".csv")
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
