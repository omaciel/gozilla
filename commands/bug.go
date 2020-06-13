package commands

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/omaciel/gozilla/bugzilla"
)

func Version() string {
	resp := bugzilla.BugzillaRequest(
		"version",
		nil,
		bugzilla.Authentication{},
	).Get()
	return resp
}

func GetBug(id string) Defects {
	url := fmt.Sprintf("bug/%v", id)
	resp := bugzilla.BugzillaRequest(
		url,
		nil,
		bugzilla.Authentication{},
	).Get()

	var defects Defects
	err := json.Unmarshal([]byte(resp), &defects)
	if err != nil {
		log.Fatalln(err)
	}
	return defects

}

func GetBugHistory(id string) BugHistoryList {
	url := fmt.Sprintf("bug/%v/history", id)
	resp := bugzilla.BugzillaRequest(
		url,
		nil,
		bugzilla.Authentication{},
	).Get()

	var history BugHistoryList
	err := json.Unmarshal([]byte(resp), &history)
	if err != nil {
		log.Fatalln(err)
	}
	return history

}
