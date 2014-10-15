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
	flagHeaders = flag.Int("headers", 0, "Number of header lines in the input file (will be repeated in each output file")
)

func main() {
	flag.Parse()
	if *flagInput == "" || *flagRecords < 1 || *flagHeaders < 0 || *flagHeaders >= *flagRecords {
		flag.Usage()
		os.Exit(1)
	}

	inputFile, err := os.Open(*flagInput)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	reader := csv.NewReader(inputFile)
	records := make([][]string, 0)
	fileCount := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		records = append(records, record)
		if len(records) == *flagRecords {
			saveCSVFile(records, fileCount)
			records = records[:*flagHeaders]
			fileCount += 1
		}
	}
	if len(records) > 0 {
		saveCSVFile(records, fileCount)
	}
}

func saveCSVFile(r [][]string, fileCount int) {
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
	writer.WriteAll(r)
}
