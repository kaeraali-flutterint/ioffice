package lib

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alicekaerast/ioffice/schema"
)

func (i *IOffice) GetRoom(search string) schema.Room {
	endpoint := fmt.Sprintf("v2/rooms/?room=%v", search)
	body := i.Request("GET", endpoint, nil)
	rooms := make([]schema.Room, 0)
	err := json.Unmarshal([]byte(body), &rooms)
	if err != nil {
		log.Fatalln(err)
	}
	if len(rooms) == 0 {
		log.Fatalf("Couldn't find any rooms for search %v", search)
	}
	return rooms[0]
}
