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

func main() {
	configure()
	buildingID := viper.GetInt("buildingID")
	roomID := viper.GetInt("roomID")

	ioffice, me := auth()

	if len(os.Args) < 2 {
		ioffice.ListReservations()
		return
	}

	switch os.Args[1] {
	case "list":
		ioffice.ListReservations()
	case "create":
		createReservation(ioffice, me, roomID, buildingID)
	case "checkin":
		checkin(ioffice)
	case "cancel":
		cancelCheckin(ioffice)
	case "buildings":
		listBuildings(ioffice)
	case "occupancy":
		listOccupancy(ioffice)
	case "floors":
		tbl := listFloors(buildingID, ioffice)
		tbl.Print()
	default:
		usage()
	}
}

func listFloors(buildingID int, ioffice *lib.IOffice) table.Table {
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
	return tbl
}

func listOccupancy(ioffice *lib.IOffice) {
	if len(os.Args) != 3 {
		fmt.Print("Must provide a floor ID for occupancy")
		fmt.Print("Usage: ioffice occupancy <floor ID>")
		os.Exit(1)
	}
	floorID, _ := strconv.Atoi(os.Args[2])
	ioffice.ShowOccupancy(floorID)
}

func listBuildings(ioffice *lib.IOffice) {
	buildings := ioffice.Buildings()

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Name")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, building := range buildings {
		tbl.AddRow(building.ID, building.Name)
	}
	tbl.Print()
}

func cancelCheckin(ioffice *lib.IOffice) {
	reservationID := os.Args[2]
	ioffice.CancelReservation(reservationID)
	ioffice.ListReservations()
}

func checkin(ioffice *lib.IOffice) {
	reservationID := os.Args[2]
	ioffice.CheckIn(reservationID)
	ioffice.ListReservations()
}

func createReservation(ioffice *lib.IOffice, me schema.User, roomID int, buildingID int) {
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
}

func usage() {
	fmt.Printf(`Please use one of the following commands:

%v list
%v create <yyyy-mm-dd> [room name]
%v checkin <reservation ID>
%v cancel <reservation ID>
%v buildings
%v floors`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}

func configure() {
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
}

func auth() (*lib.IOffice, schema.User) {
	ioffice := lib.NewIOffice(viper.GetString("hostname"), viper.GetString("username"), viper.GetString("password"), viper.GetString("session"))
	me := ioffice.GetMe()
	if !ioffice.WasOkay() {
		log.Fatalln("Stopping now as auth failed.  Are you on SSO?  See README.md on how to authenticate.")
	}
	return ioffice, me
}
