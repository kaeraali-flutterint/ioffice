package lib

import (
	"encoding/json"

	"github.com/alicekaerast/ioffice/schema"
)

func (i *IOffice) GetMe() schema.User {
	endpoint := "v2/users/me"
	body := i.Request("GET", endpoint, nil)
	user := schema.User{}
	json.Unmarshal([]byte(body), &user)
	return user
}
