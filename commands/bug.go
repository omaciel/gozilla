package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	"github.com/omaciel/gozilla/bugzilla"
)

// Version : Fetch the version information of the Bugzilla server.
func Version() string {
	resp := bugzilla.BugzillaRequest(
		"version",
		nil,
		bugzilla.Authentication{},
	).Get()
	return resp
}

// GetBug : Fetch a bug/s matching an id or a set of keywords. Also supports adding a limit of results.
func GetBug(filter, limit string) Defects {
	var url string

	if matched, _ := regexp.MatchString("^[0-9]+$", filter); matched {
		url = fmt.Sprintf("bug/%v", filter)
	} else {
		url = fmt.Sprintf("bug?keywords=%v&limit=%v", filter, limit)
	}

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

// GetBugHistory : Fetch the history of a bug.
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
