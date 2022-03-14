package lib

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type IOffice struct {
	hostname   string
	username   string
	password   string
	httpClient *http.Client
	lastStatus int
}

func NewIOffice(hostname string, username string, password string) *IOffice {
	return &IOffice{
		hostname:   hostname,
		username:   username,
		password:   password,
		lastStatus: 0,
	}
}

func (i *IOffice) WasOkay() bool {
	return (i.lastStatus > 199 && i.lastStatus < 300)
}

func (i *IOffice) Request(method string, endpoint string, body io.Reader) []byte {
	url := fmt.Sprintf("https://%v.iofficeconnect.com/external/api/rest/%v", i.hostname, endpoint)
	req, err := http.NewRequest(method, url, body)
	if method == "POST" {
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	req.Header.Add("x-auth-username", i.username)
	req.Header.Add("x-auth-password", i.password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	i.lastStatus = resp.StatusCode
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		fmt.Println("Error status detected: ", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	return respBody
}
