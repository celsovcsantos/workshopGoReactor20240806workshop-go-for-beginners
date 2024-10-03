package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {

	useCsvFile := flag.String("file", "testdata/users_001.csv", "Path to file")
	nWinners := flag.Int("winners", 6, "Number of winners")
	records, err := ParseCsvFile("testdata/users_001.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	// for _, record := range records {
	// 	fmt.Println(record)
	// }
	winners := SelectWinners(records, 3)
	for _, winner := range winners {
		fmt.Println(winner)
	}
}

func SelectWinners(records [][]string, nWinners int) [][]string {
	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})
	return records[:nWinners]
}
