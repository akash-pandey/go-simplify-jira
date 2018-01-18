package common

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

// ReadFile reads all recorrds from a CSV file, Illegal records are  be ignored
func ReadFile(fileName string) (record [][]string, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(bufio.NewReader(f))
	records := [][]string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			continue
		}
		records = append(records, record)
	}
	return records, err
}
