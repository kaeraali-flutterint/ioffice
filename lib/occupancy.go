package lib

import (
	"encoding/json"
	"fmt"
	"github.com/alicekaerast/ioffice/schema"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"log"
	"sort"
	"time"
)

func (i *IOffice) ShowOccupancy(floor int) {
	occupancy := i.GetOccupancy(floor)

	reservations := make([]schema.Reservation, 0)
	for _, roomReservations := range occupancy {
		for _, reservation := range roomReservations.Reservations {
			reservation.Room = schema.Room{Name: roomReservations.Name, ID: roomReservations.ID}
			reservations = append(reservations, reservation)
		}
	}

	sort.Slice(reservations, func(i, j int) bool {
		iStartdate := time.Unix(reservations[i].StartDate/1000, 0).Format("2006-01-02")
		jStartdate := time.Unix(reservations[j].StartDate/1000, 0).Format("2006-01-02")

		//return reservations[i].StartDate < reservations[j].StartDate
		if iStartdate != jStartdate {
			return iStartdate < jStartdate
		}
		return reservations[i].Room.Name < reservations[j].Room.Name
	})

	fmt.Printf("Occupancy for floor ID %v: %v\n", floor, len(occupancy))
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("Location", "Start", "Who")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, r := range reservations {
		unixTimeUTC := time.Unix(r.StartDate/1000, 0)
		tbl.AddRow(r.Room.Name, unixTimeUTC.Format(time.RFC822), r.User.Name)
	}

	tbl.Print()
}

func (i *IOffice) GetOccupancy(floor int) []schema.RoomReservations {
	endpoint := fmt.Sprintf("v2/rooms/?floorId=%v", floor)
	endpoint = endpoint + fmt.Sprintf("&selector=anonymousReservations(endDate,numberOfPeople,startDate,user)")
	endpoint = endpoint + fmt.Sprintf("&includeNonReservable=false&includeReservable=true&limit=1000")

	body := i.Request("GET", endpoint, nil)
	rooms := make([]schema.RoomReservations, 0)
	err := json.Unmarshal([]byte(body), &rooms)

	if err != nil {
		log.Fatalln(err)
	}

	return rooms
}
