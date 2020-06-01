package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Defects struct {
	Bugs []Defect `json:"bugs"`
}

type Defect struct {
	ID       int    `json:"id"`
	Product  string `json:"product"`
	Summary  string `json:"summary"`
	Severity string `json:"severity"`
}

var client = &http.Client{}

func execute(req http.Request) string {
	resp, err := client.Do(&req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	s := string(bodyText)

	return s
}

func Get(url string) http.Request {
	req, _ := http.NewRequest(
		"GET",
		url,
		nil)
	//req.SetBasicAuth(username, passwd)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	log.Print(req.URL.String())
	return *req
}

func Version() string {
	url := fmt.Sprintf("%v/version", BugzillaURL)
	resp := execute(Get(url))
	return resp
}

func Bug(id string) Defects {
	url := fmt.Sprintf("%v/bug/%v", BugzillaURL, id)
	resp := execute(Get(url))

	var defects Defects
	err := json.Unmarshal([]byte(resp), &defects)
	if err != nil {
		log.Fatalln(err)
	}
	return defects

}
