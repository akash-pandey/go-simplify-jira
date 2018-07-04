package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	jira "github.com/andygrunwald/go-jira"
	"gitlab.com/akash.pandey/go-simplify-jira/csv"
	"golang.org/x/crypto/ssh/terminal"
)

// TODO: accept command line configs as well
func main() {

	var url string
	var username string
	var path string
	fmt.Printf("Path to csv file: ")
	fmt.Scanln(&path)
	fmt.Printf("Enter Jira URL: ")
	fmt.Scanln(&url)
	fmt.Printf("Enter Username: ")
	fmt.Scanln(&username)
	fmt.Printf("Enter Password: ")
	passwd, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(passwd)
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}

	jiraClient, err := jira.NewClient(tp.Client(), url)
	if err != nil {
		panic(err)
	}

	// todo: As a foundation to batch operations, create a new function to return rows grouped by action type.
	records, err := parseFile(path)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	for _, record := range records {
		var response *jira.Issue
		var actionStr string
		switch record.action {
		case create:
			response, err = createIssue(record.issue, jiraClient)
			actionStr = "Created"
		case get:
			response, err = getIssueByKey(record.issue.Key, jiraClient)
			actionStr = "Fetched"
		default:
			response, err = nil, fmt.Errorf("Invalid TransactionType: %s", record.action.string())
		}
		if err != nil {
			fmt.Println(err)
			payload, _ := record.issue.Fields.MarshalJSON()
			fmt.Printf("Payload: %s - %s \n %s\n", record.issue.Self, record.issue.Key, string(payload[:]))
		} else {
			fmt.Printf("Successfully %s issue. Key: %s\n", actionStr, response.Key)
		}
	}
}

func createIssue(record *jira.Issue, jiraClient *jira.Client) (issue *jira.Issue, err error) {
	parent, err1 := getIssueByKey(record.Fields.Parent.Key, jiraClient)
	if err1 != nil {
		return nil, fmt.Errorf("No parent ticket exist with key: %s", record.Fields.Parent.Key)
	}
	record.Fields.Parent.ID = parent.ID
	issue, _, err = jiraClient.Issue.Create(record)
	return issue, err
}

func updateIssue(record []string, jiraClient *jira.Client) (issue *jira.Issue, err error) {
	issue, _, err = jiraClient.Issue.Get("MCPU-6640", nil)
	return issue, err

}

func queryIssues(record issueWrapper, jiraClient *jira.Client) (issue *jira.Issue, err error) {
	issue, _, err = jiraClient.Issue.Get("MCPU-6640", nil)
	return issue, err
}

func getIssueByKey(key string, jiraClient *jira.Client) (issue *jira.Issue, err error) {
	issue, _, err = jiraClient.Issue.Get(key, nil)
	return issue, err
}

func parseFile(fileName string) (issues []issueWrapper, err error) {
	fmt.Printf("\nReading file: %s\n", fileName)
	records, err := csv.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Total %d records found\n", len(records))
	issues = []issueWrapper{}
	for _, record := range records {
		action := getAction(record[0])
		if action != invalid {
			issue := issueWrapper{action: action, issue: getIssueFields(action, record)}
			issues = append(issues, issue)
		} else {
			fmt.Printf("WARN: Invalid Jira operation provided: %s.\n", record[0])
		}
	}
	fmt.Printf("Total %d valid records found\n", len(issues))
	return issues, err
}

func getAction(action string) jiraAction {
	for index, val := range AllowedActions {
		if action == val {
			return jiraAction(index)
		}
	}
	return invalid
}

func getIssueFields(action jiraAction, record []string) *jira.Issue {
	switch action {
	case create:
		return getIssueFieldsForCreate(record)
	case update:
	case get:
		return getIssueFieldsForUpdate(record)
	}
	return nil
}

func getIssueFieldsForUpdate(record []string) *jira.Issue {
	return &jira.Issue{
		Key: record[3],
		Fields: &jira.IssueFields{
			Project: jira.Project{
				Key: record[2],
			},
			Parent: &jira.Parent{
				Key: record[4],
			},
			Type: jira.IssueType{
				Name: record[1],
			},
			Summary:     record[5],
			Description: record[6],
			Assignee: &jira.User{
				Name: record[7],
			},
		},
	}
}

func getIssueFieldsForCreate(record []string) *jira.Issue {
	const format = "2006-Jan-02" // TODO - parameterize
	dueDate, err := time.Parse(format, record[9])
	if err != nil {
		log.Fatal(err)
	}
	return &jira.Issue{
		Fields: &jira.IssueFields{
			Type: jira.IssueType{
				Name:    record[1],
				Subtask: record[1] == "Sub-task",
			},
			Project: jira.Project{
				Key: record[2],
			},
			Parent: &jira.Parent{
				Key: record[4],
			},
			Summary:     record[5],
			Description: record[6],
			Assignee: &jira.User{
				Name: record[7],
			},
			TimeTracking: &jira.TimeTracking{
				OriginalEstimate: record[8],
			},
			Duedate: jira.Date(dueDate),
		},
	}
}
