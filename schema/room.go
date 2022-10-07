package schema

type Room struct {
	Area            float64 `json:"area"`
	LongDescription string  `json:"longDescription"`
	Reservable      bool    `json:"reservable"`
	Description     string  `json:"description"`
	Type            struct {
		HexColor    string  `json:"hexColor"`
		ContentFlag int     `json:"contentFlag"`
		Cost        float64 `json:"cost"`
		DateCreated int64   `json:"dateCreated"`
		Color       struct {
		} `json:"color"`
		Name         string `json:"name"`
		Occupiable   bool   `json:"occupiable"`
		ID           int    `json:"id"`
		ParkingSpace bool   `json:"parkingSpace"`
		DateUpdated  int64  `json:"dateUpdated"`
		TypeCode     string `json:"typeCode"`
	} `json:"type"`
	Capacity    int   `json:"capacity"`
	DateUpdated int64 `json:"dateUpdated"`
	DateCreated int64 `json:"dateCreated"`
	GoogleData  struct {
		GoogleCalResourceID string `json:"googleCalResourceId"`
		GoogleCalAddress    string `json:"googleCalAddress"`
	} `json:"googleData,omitempty"`
	ReservableByRules bool   `json:"reservableByRules"`
	Name              string `json:"name"`
	RemoteInfo        string `json:"remoteInfo"`
	ID                int    `json:"id"`
	Floor             Floor  `json:"floor"`
	Occupied          bool   `json:"occupied"`
}

type Reservation struct {
	EndDate int64 `json:"endDate"`
	ID      int   `json:"id"`
	User    struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"user"`
	NumberOfPeople int   `json:"numberOfPeople"`
	StartDate      int64 `json:"startDate"`
	Room           Room
}

type RoomReservations struct {
	Name         string        `json:"name"`
	Reservations []Reservation `json:"anonymousReservations"`
	ID           int           `json:"id"`
}
