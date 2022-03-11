package schema

type Reservations []struct {
	CheckedOut         bool   `json:"checkedOut"`
	Notes              string `json:"notes"`
	CancellationReason string `json:"cancellationReason"`
	CheckedIn          bool   `json:"checkedIn"`
	EndDate            int64  `json:"endDate"`
	Center             struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"center"`
	Room           Room   `json:"room"`
	DateCreated    int64  `json:"dateCreated"`
	Name           string `json:"name"`
	Cancelled      bool   `json:"cancelled"`
	ID             int    `json:"id"`
	User           User   `json:"user"`
	NumberOfPeople int    `json:"numberOfPeople"`
	StartDate      int64  `json:"startDate"`
}

type ReservationRequest struct {
	Guests []interface{} `json:"guests"`
	Notes  string        `json:"notes"`
	User   struct {
		ID int `json:"id"`
	} `json:"user"`
	Center struct {
		ID int `json:"id"`
	} `json:"center"`
	Room struct {
		ID int `json:"id"`
	} `json:"room"`
	NumberOfPeople int   `json:"numberOfPeople"`
	StartDate      int64 `json:"startDate"`
	EndDate        int64 `json:"endDate"`
	AllDay         bool  `json:"allDay"`
}
