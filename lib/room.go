package lib

import (
	"encoding/json"
	"fmt"
	"github.com/alicekaerast/ioffice/schema"
	"io/ioutil"
	"log"
	"net/http"
)

func GetRoom(username string, password string, hostname string, search string) schema.Room {
	url := fmt.Sprintf("https://%v.iofficeconnect.com/external/api/rest/v2/rooms/?room=%v", hostname, search)
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
	rooms := make([]schema.Room, 0)
	json.Unmarshal([]byte(body), &rooms)

	fmt.Println(rooms)
	if len(rooms) == 0 {
		log.Fatalf("Couldn't find any rooms for search %v", search)
	}

	return rooms[0]
}
