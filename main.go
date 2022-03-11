package main

import (
	"fmt"
	"github.com/alicekaerast/ioffice/lib"
	"github.com/araddon/dateparse"
	"github.com/spf13/viper"
	"os"
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

	me := lib.GetMe(username, password, hostname)

	if len(os.Args) == 2 {
		lib.CreateReservation(username, password, hostname, me, roomID, dateparse.MustParse(os.Args[1]))
	}

	if len(os.Args) == 3 {
		room := lib.GetRoom(username, password, hostname, os.Args[2])
		lib.CreateReservation(username, password, hostname, me, room.ID, dateparse.MustParse(os.Args[1]))
	}

	lib.ListReservations(username, password, hostname)
}
