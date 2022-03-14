package lib

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
)

type IOffice struct {
	hostname   string
	username   string
	password   string
	session    string
	httpClient *http.Client
	lastStatus int
}

func NewIOffice(hostname string, username string, password string, session string) *IOffice {
	jar, _ := cookiejar.New(nil)
	return &IOffice{
		hostname:   hostname,
		username:   username,
		password:   password,
		session:    session,
		lastStatus: 0,
		httpClient: &http.Client{
			Jar: jar,
		},
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
	if i.session == "" {
		req.Header.Add("x-auth-username", i.username)
		req.Header.Add("x-auth-password", i.password)
	} else {
		req.AddCookie(&http.Cookie{
			Name:  "ACTID",
			Value: i.session,
		})
	}
	resp, err := i.httpClient.Do(req)
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
