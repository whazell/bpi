package bpi_test

import (
	bba "bpi"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	/*	players, err := bba.Players("2018")

		if err != nil {
			fmt.Println(err)
			return
		}

		currys := bba.FilterPlayers(players, func(p bba.Player) bool {
			return p.LastName == "Curry"
		})

		for _, i := range currys {
			fmt.Println(i)
		}

	*/

	/*
		teams, _ := bba.NBATeams("2018")
		for _, i := range teams {
			fmt.Println(i)
		}
	*/

	scoreboards, _ := bba.Scoreboards("20181107")
	fmt.Println(scoreboards)
}