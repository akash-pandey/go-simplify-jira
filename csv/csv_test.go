package csv_test

import (
	"fmt"

	"com.gitlab/akash.pandey/go-simplify-jira/csv"
)

func ExampleReadFile() {
	fmt.Println(csv.ReadFile("c:\\test\\test_csv.csv"))
	// Output: [[New  MCPU-AAAA Test Ticket Sample Body of Test Ticket akash.pandey 1d ] [New  MCPU-BBBB Test Ticket Sample Body of Test Ticket akash.pandey 1d ] [MODIFY MCPU-AAAA      {assignee: rajesh.bhatt, originalEstimates: 2d}]] <nil>
}
