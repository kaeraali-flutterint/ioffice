package lib

import (
	"encoding/json"
	"fmt"
	"github.com/alicekaerast/ioffice/schema"
	"io/ioutil"
	"log"
	"net/http"
)

func GetRoom(username string, password string, hostname string, ID string) schema.Room {
	url := fmt.Sprintf("https://%v.iofficeconnect.com/external/api/rest/v2/rooms/%v", hostname, ID)
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

	room := schema.Room{}
	json.Unmarshal([]byte(body), &room)
	return room
}
