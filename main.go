package main

import (
	"fmt"
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

	ioffice := lib.NewIOffice(hostname, username, password)

	me := ioffice.GetMe()

	if len(os.Args) == 2 {
		ioffice.CreateReservation(me, roomID, dateparse.MustParse(os.Args[1]))
	}

	ioffice.ListReservations()
}
