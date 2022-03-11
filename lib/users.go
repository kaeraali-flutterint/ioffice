package lib

import (
	"encoding/json"
	"fmt"
	"github.com/alicekaerast/ioffice/schema"
	"io/ioutil"
	"log"
	"net/http"
)

func GetMe(username string, password string, hostname string) schema.User {
	url := fmt.Sprintf("https://%v.iofficeconnect.com/external/api/rest/v2/users/me", hostname)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("x-auth-username", username)
	req.Header.Add("x-auth-password", password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	user := schema.User{}
	json.Unmarshal([]byte(body), &user)
	return user
}
