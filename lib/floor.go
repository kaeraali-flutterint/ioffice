package lib

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alicekaerast/ioffice/schema"
)

func (i *IOffice) Floors() []schema.Floor {
	endpoint := fmt.Sprintf("v2/floors")
	body := i.Request("GET", endpoint, nil)
	floors := make([]schema.Floor, 0)
	err := json.Unmarshal([]byte(body), &floors)
	if err != nil {
		log.Fatalln(err)
	}
	return floors
}

func (i *IOffice) FloorsForBuilding(building string) []schema.Floor {
	endpoint := fmt.Sprintf("v2/floors")
	endpoint = endpoint + fmt.Sprintf("?buildingId=%v", building)
	body := i.Request("GET", endpoint, nil)
	floors := make([]schema.Floor, 0)
	err := json.Unmarshal([]byte(body), &floors)
	if err != nil {
		log.Fatalln(err)
	}
	return floors
}
