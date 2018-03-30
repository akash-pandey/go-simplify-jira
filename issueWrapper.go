package main

import jira "github.com/andygrunwald/go-jira"

// AllowedActions is list of jira operations allowed through go-simplify-jira, internally representd by enum jiraAction
var AllowedActions = [...]string{
	"INVALID",
	"GET",
	"CREATE",
	"UPDATE",
	"DELETE",
	"QUERY",
}

type jiraAction int

const (
	invalid jiraAction = iota
	get
	create
	update
	delete
	query
)

func (action jiraAction) string() string {
	if action > query {
		action = invalid
	}
	return AllowedActions[action]
}

type issueWrapper struct {
	action jiraAction
	issue  *jira.Issue
}
