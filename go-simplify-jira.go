package main

import (
	"fmt"
	"log"

	"com.gitlab/akash.pandey/go-simplify-jira/csv"
)

func main() {
	// todo: As a foundation to batch operations, create a new function to return rows grouped by action type.
	records, err := csv.ReadFile("c:\\test\\test_csv.csv")
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		fmt.Println(record)
	}

}
