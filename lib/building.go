package lib

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alicekaerast/ioffice/schema"
)

func (i *IOffice) Buildings() []schema.Building {
	endpoint := fmt.Sprintf("v2/buildings")
	body := i.Request("GET", endpoint, nil)
	buildings := make([]schema.Building, 0)
	err := json.Unmarshal([]byte(body), &buildings)
	if err != nil {
		log.Fatalln(err)
	}
	return buildings
}
