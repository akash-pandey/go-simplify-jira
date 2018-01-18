package main

import (
	"fmt"
	"log"

	"gitlab.com/akash.pandey/go-simplify-jira/common"
)

func main() {
	records, err := common.ReadFile("c:\\test\\test_csv.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(records)
}
