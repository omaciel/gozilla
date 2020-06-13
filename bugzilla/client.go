package bugzilla

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/omaciel/gozilla/commands"
)

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

type Authentication struct {
	Username string
	Password string
}

type Request struct {
	URL      string
	EndPoint string
	Auth     Authentication
	Data     map[string]string
}

func (r Request) Get() string {
	req, _ := http.NewRequest(
		"GET",
		r.URL,
		nil)
	//req.SetBasicAuth(username, passwd)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	log.Print(req.URL.String())
	return execute(*req)
}

func BugzillaRequest(
	endpoint string,
	data map[string]string,
	auth Authentication) Request {
	request := Request{}
	request.URL = fmt.Sprintf("%v/%v", commands.BugzillaURL, endpoint)
	request.Auth = auth
	request.Data = data
	return request
}
