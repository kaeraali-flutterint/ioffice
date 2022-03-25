package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/alicekaerast/ioffice/schema"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func (i *IOffice) ShowOccupancy(floor int) {
	occupancy := i.GetOccupancy(floor)
	fmt.Printf("Occupancy: %v\n", len(occupancy))

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("Location", "Start", "Who")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, o := range occupancy {
		if len(o.Reservations) > 0 {
			for _, reservation := range o.Reservations {
				unixTimeUTC := time.Unix(reservation.StartDate/1000, 0)
				tbl.AddRow(o.Name, unixTimeUTC.Format(time.RFC822), reservation.User.Name)
			}
		}
	}
	tbl.Print()
}

func (i *IOffice) GetOccupancy(floor int) []schema.RoomReservations {
	endpoint := fmt.Sprintf("v2/rooms/?floorId=%v", floor)
	endpoint = endpoint + fmt.Sprintf("&selector=anonymousReservations(endDate,numberOfPeople,startDate,user)")
	endpoint = endpoint + fmt.Sprintf("&includeNonReservable=false&includeReservable=true&limit=1000")

	fmt.Printf("Request: %v\n", endpoint)

	body := i.Request("GET", endpoint, nil)
	rooms := make([]schema.RoomReservations, 0)
	err := json.Unmarshal([]byte(body), &rooms)

	if err != nil {
		log.Fatalln(err)
	}

	return rooms
}
