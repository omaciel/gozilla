package commands

import (
    "io/ioutil"
    "net/http"
    "log"
    "fmt"
)

type Defect struct {
    ID string `json:"id"`
    Product string `json:"product"`
    Summary string `json:"summary"`
    Severity string `json:"severity"`
}
var client = &http.Client{}
	
func execute(req http.Request) string {
    resp, err := client.Do(&req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)
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
    return execute(Get(url))
}

func Bug(id string) string {
    url := fmt.Sprintf("%v/bug/%v", BugzillaURL, id)
    return execute(Get(url))
}

