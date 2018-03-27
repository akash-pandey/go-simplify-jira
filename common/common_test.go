package common_test

import (
	"fmt"
	"gitlab.com/akash.pandey/go-simplify-jira/common"
)

func ExampleReadFile()  {
	fmt.Println(common.ReadFile("c:\\test\\test_csv.csv"))
	// Output: [[akash pandey gurgaon 9911900821] [tanvi pathak gurgaon 9911900821]] <nil>
}