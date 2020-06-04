package commands

import (
	"io/ioutil"
	"log"
	"net/http"
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
