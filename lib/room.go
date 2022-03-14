package lib

import (
	"encoding/json"
	"fmt"

	"github.com/alicekaerast/ioffice/schema"
)

func (i *IOffice) GetRoom(ID string) schema.Room {
	endpoint := fmt.Sprintf("v2/rooms/%v", ID)
	body := i.Request("GET", endpoint, nil)
	room := schema.Room{}
	json.Unmarshal([]byte(body), &room)
	return room
}
