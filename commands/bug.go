package commands

import (
	"encoding/json"
	"fmt"
	"log"
)

func Version() string {
	url := fmt.Sprintf("%v/version", BugzillaURL)
	resp := execute(Get(url))
	return resp
}

func GetBug(id string) Defects {
	url := fmt.Sprintf(BugByID, BugzillaURL, id)
	resp := execute(Get(url))

	var defects Defects
	err := json.Unmarshal([]byte(resp), &defects)
	if err != nil {
		log.Fatalln(err)
	}
	return defects

}

func GetBugHistory(id string) BugHistoryList {
	url := fmt.Sprintf(BugHistoryByID, BugzillaURL, id)
	resp := execute(Get(url))

	var history BugHistoryList
	err := json.Unmarshal([]byte(resp), &history)
	if err != nil {
		log.Fatalln(err)
	}
	return history

}
