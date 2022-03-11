package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alicekaerast/ioffice/schema"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ListReservations(username string, password string, hostname string) {
	url := fmt.Sprintf("https://%v.iofficeconnect.com/external/api/rest/v2/reservations/?showOnlyMyReservations=true", hostname)
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

	reservations := schema.Reservations{}
	json.Unmarshal([]byte(body), &reservations)
	log.Printf("Upcoming reservations: %v", len(reservations))

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Start", "Location Name", "Location ID")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, reservation := range reservations {
		unixTimeUTC := time.Unix(reservation.StartDate/1000, 0)
		tbl.AddRow(reservation.ID, unixTimeUTC.Format(time.RFC822), reservation.Room.Name, reservation.Room.ID)
	}
	tbl.Print()
}

func CreateReservation(username string, password string, hostname string, user schema.User, roomID int, date time.Time) {
	url := fmt.Sprintf("https://%v.iofficeconnect.com/external/api/rest/v2/reservations", hostname)

	// {"guests":[],"notes":"","user":{"id":2409},"center":{"id":74},"room":{"id":672},"numberOfPeople":1,"startDate":1647011700000,"endDate":1647015300000,"allDay":false}
	reservationRequest := schema.ReservationRequest{
		Guests: nil,
		Notes:  "",
		User: struct {
			ID int `json:"id"`
		}{
			ID: user.ID,
		},
		Center: struct {
			ID int `json:"id"`
		}{
			ID: 74,
		},
		Room: struct {
			ID int `json:"id"`
		}{
			ID: roomID,
		},
		NumberOfPeople: 1,
		StartDate:      date.Unix() * 1000,
		EndDate:        date.Unix() * 1000,
		AllDay:         true,
	}

	jsonReservationRequest, _ := json.Marshal(reservationRequest)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReservationRequest))
	req.Header.Add("x-auth-username", username)
	req.Header.Add("x-auth-password", password)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
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

	fmt.Println(string(body))

}
