package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alicekaerast/ioffice/lib"
	"github.com/araddon/dateparse"
	"github.com/spf13/viper"
)

func main() {
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
	roomID := viper.GetInt("roomID")
	hostname := viper.GetString("hostname")
	session := viper.GetString("session")

	ioffice := lib.NewIOffice(hostname, username, password, session)

	me := ioffice.GetMe()
	if !ioffice.WasOkay() {
		log.Println("Stopping now as auth failed.  Are you on SSO?  See README.md on how to authenticate.")
		return
	}

	if len(os.Args) == 2 {
		ioffice.CreateReservation(me, roomID, dateparse.MustParse(os.Args[1]))
	}

	if len(os.Args) == 3 {
		room := ioffice.GetRoom(os.Args[2])
		ioffice.CreateReservation(me, room.ID, dateparse.MustParse(os.Args[1]))
	}

	ioffice.ListReservations()
}
