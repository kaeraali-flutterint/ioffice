package main

import (
	"fmt"
	"github.com/alicekaerast/ioffice/lib"
	"github.com/alicekaerast/ioffice/schema"
	"github.com/araddon/dateparse"
	"github.com/fatih/color"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/rodaine/table"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"path"
)

var k = koanf.New(".")

func main() {
	configure()
	buildingID := k.Int("buildingID")
	roomID := k.Int("roomID")

	ioffice, me := auth()

	app := &cli.App{
		Suggest: true,
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "list my reservations",
				Action: func(cCtx *cli.Context) error {
					ioffice.ListReservations()
					return nil
				},
			},
			{
				Name:  "create",
				Usage: "create a reservation for a date provided as an argument",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "location",
						Value: "0",
						Usage: "location to reserve",
					},
					&cli.IntFlag{
						Name:  "building",
						Value: buildingID,
						Usage: "building ID within which to search",
					},
				},
				Action: func(cCtx *cli.Context) error {
					if cCtx.Int("building") != 0 && cCtx.String("location") != "0" {
						roomID = ioffice.GetRoomWithBuilding(cCtx.String("location"), cCtx.Int("building")).ID
					} else if cCtx.String("location") != "0" {
						roomID = ioffice.GetRoom(cCtx.String("location")).ID
					}

					if cCtx.NArg() > 0 {
						ioffice.CreateReservation(me, roomID, dateparse.MustParse(cCtx.Args().Get(0)))
					} else {
						return cli.Exit("requires a date to reserve for. Usage: ioffice create 2023-10-31", 1)
					}
					ioffice.ListReservations()
					return nil
				},
			},
			{
				Name:  "checkin",
				Usage: "checkin to a reservation",
				Action: func(cCtx *cli.Context) error {
					if cCtx.NArg() == 1 {
						reservationID := cCtx.Args().Get(0)
						ioffice.CheckIn(reservationID)
						ioffice.ListReservations()
						return nil
					} else {
						return cli.Exit("requires a reservation ID to checkin to", 1)
					}
				},
			},
			{
				Name:  "cancel",
				Usage: "cancel a reservation",
				Action: func(cCtx *cli.Context) error {
					if cCtx.NArg() == 1 {
						reservationID := cCtx.Args().Get(0)
						ioffice.CancelReservation(reservationID)
						ioffice.ListReservations()
						return nil
					} else {
						return cli.Exit("requires a reservation ID to cancel", 1)
					}
				},
			},
			{
				Name:  "buildings",
				Usage: "list buildings",
				Action: func(cCtx *cli.Context) error {
					buildings := ioffice.Buildings()

					headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
					columnFmt := color.New(color.FgYellow).SprintfFunc()
					tbl := table.New("ID", "Name")
					tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

					for _, building := range buildings {
						tbl.AddRow(building.ID, building.Name)
					}
					tbl.Print()
					return nil
				},
			},
			{
				Name:  "occupancy",
				Usage: "show occupancy for a floor",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "floor",
						Usage:    "floor ID",
						Required: true,
					},
				},
				Action: func(cCtx *cli.Context) error {
					ioffice.ShowOccupancy(cCtx.Int("floor"))
					return nil
				},
			},
			{
				Name:  "floors",
				Usage: "show floors in a building",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "building",
						Usage:   "building ID",
						Aliases: []string{"b"},
						Value:   buildingID,
					},
				},
				Action: func(cCtx *cli.Context) error {
					buildingID = cCtx.Int("building")
					buildings := ioffice.Buildings()
					idx := slices.IndexFunc(buildings, func(c schema.Building) bool { return c.ID == buildingID })

					fmt.Printf("Floors for %v\n", buildings[idx].Name)
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
					return nil
				},
			},
		},
		Usage: "manage ioffice reservations",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func configure() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	k.Load(confmap.Provider(map[string]interface{}{"buildingID": 0}, "."), nil)

	k.Load(file.Provider(path.Join(dirname, "ioffice.yaml")), yaml.Parser())
	k.Load(file.Provider(path.Join(dirname, ".config", "ioffice.yaml")), yaml.Parser())
	k.Load(file.Provider("ioffice.yaml"), yaml.Parser())
}

func auth() (*lib.IOffice, schema.User) {
	ioffice := lib.NewIOffice(k.String("hostname"), k.String("username"), k.String("password"), k.String("session"))
	me := ioffice.GetMe()
	if !ioffice.WasOkay() {
		log.Fatalln("Stopping now as auth failed.  Are you on SSO?  See README.md on how to authenticate.")
	}
	return ioffice, me
}
