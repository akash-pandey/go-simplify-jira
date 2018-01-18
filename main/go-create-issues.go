package main

func main() {
	records, err := common.readFile("c:\\test\\test_csv.csv")
	if(err != nil) {
		log.fatal(err)
	}
	fmt.println(records)
}
