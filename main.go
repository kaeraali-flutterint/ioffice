package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/alicekaerast/ioffice/lib"
	"github.com/alicekaerast/ioffice/schema"
	"github.com/araddon/dateparse"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/viper"
)

func usage() {
	fmt.Printf(`Please use one of the following commands:

%v list
%v create <yyyy-mm-dd> [room name]
%v checkin <reservation ID>
%v cancel <reservation ID>
%v buildings`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}

func main() {
	viper.SetDefault("buildingID", 0)
	viper.SetConfigName("ioffice")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	viper.SetEnvPrefix("ioffice")
	viper.AutomaticEnv()

	username := viper.GetString("username")
	password := viper.GetString("password")
	buildingID := viper.GetInt("buildingID")
	roomID := viper.GetInt("roomID")
	hostname := viper.GetString("hostname")
	session := viper.GetString("session")

	ioffice := lib.NewIOffice(hostname, username, password, session)

	me := ioffice.GetMe()
	if !ioffice.WasOkay() {
		log.Println("Stopping now as auth failed.  Are you on SSO?  See README.md on how to authenticate.")
		return
	}

	if len(os.Args) < 2 {
		ioffice.ListReservations()
	} else {

		switch os.Args[1] {
		case "list":
			ioffice.ListReservations()
		case "create":
			if len(os.Args) == 2 {
				usage()
			}
			if len(os.Args) == 3 {
				ioffice.CreateReservation(me, roomID, dateparse.MustParse(os.Args[2]))
			}
			if len(os.Args) == 4 {
				room := schema.Room{}
				if buildingID == 0 {
					room = ioffice.GetRoom(os.Args[3])
				} else {
					room = ioffice.GetRoomWithBuilding(os.Args[3], buildingID)
				}
				ioffice.CreateReservation(me, room.ID, dateparse.MustParse(os.Args[2]))
			}
			ioffice.ListReservations()
		case "checkin":
			reservationID := os.Args[2]
			ioffice.CheckIn(reservationID)
			ioffice.ListReservations()
		case "cancel":
			reservationID := os.Args[2]
			ioffice.CancelReservation(reservationID)
			ioffice.ListReservations()
		case "buildings":
			buildings := ioffice.Buildings()

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()
			tbl := table.New("ID", "Name")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, building := range buildings {
				tbl.AddRow(building.ID, building.Name)
			}
			tbl.Print()
		case "occupancy":
			if len(os.Args) != 3 {
				fmt.Print("Must provide a floor ID for occupancy")
				fmt.Print("Usage: ioffice occupancy <floor ID>")
				os.Exit(1)
			}
			floorID, _ := strconv.Atoi(os.Args[2])
			ioffice.ShowOccupancy(floorID)
		case "floors":
			if len(os.Args) == 3 {
				buildingID, _ = strconv.Atoi(os.Args[2])
			}
			floors := make([]schema.Floor, 0)
			if buildingID == 0 {
				floors = ioffice.Floors()
			} else {
				floors = ioffice.FloorsForBuilding(fmt.Sprint(buildingID))
			}

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()
			tbl := table.New("ID", "Name")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, floor := range floors {
				tbl.AddRow(floor.ID, floor.Name)
			}
			tbl.Print()

		default:
			usage()
		}
	}
}
