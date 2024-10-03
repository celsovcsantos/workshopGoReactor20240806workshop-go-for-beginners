package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ParseCsvFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ParseCsvFile - error opening csv file: %w", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("ParseCsvFile - error reading records: %w", err)
	}
	return records, nil
}
