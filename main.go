// csvsplit: Split a .csv into multiple files.
// https://github.com/JeffPaine/csvsplit
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
	flagInput   = flag.String("input", "", "Filename of the input file to split (if blank, uses stdin)")
	flagOutput  = flag.String("output", "", "filename / path of the file output (optional)")
	flagHeaders = flag.Int("headers", 0, "Number of header lines in the input file (will be repeated in each output file")
)

func main() {
	flag.Parse()

	// Sanity check flags
	if *flagRecords < 1 || *flagHeaders < 0 || *flagHeaders >= *flagRecords {
		flag.Usage()
		os.Exit(1)
	}

	// Get input from a given file or stdin
	var r *csv.Reader
	if *flagInput != "" {
		f, err := os.Open(*flagInput)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		r = csv.NewReader(f)
	} else {
		r = csv.NewReader(os.Stdin)
	}

	var recs [][]string
	count := 1
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		recs = append(recs, record)
		if len(recs) == *flagRecords {
			save(&recs, count)
			// Reset records to include just the header lines (if any)
			recs = recs[:*flagHeaders]
			count++
		}
	}
	if len(recs) > 0 {
		save(&recs, count)
	}
}

// save saves the given *[][]string of csv data to a .csv file. Files are named
// sequentially in the form of 001.csv, 002.csv, etc.
func save(recs *[][]string, c int) {
	name := fmt.Sprintf("%v%03d%v", *flagOutput, c, ".csv")

	// Make sure we don't overwrite existing files
	if _, err := os.Stat(name); err == nil {
		log.Fatal("File exists: ", name)
	}

	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteAll(*recs)
}
