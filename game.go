package bpi

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strconv"
	"time"
)

type Game struct {
	League            string
	GameID            string
	SeasonStageID     int
	GameURLCode       string
	StatusNum         int
	ExtendedStatusNum int
	StartTimeUTC      string
	StartDateEastern  time.Time
	StartTimeEastern  string
	IsBuzzerBeater    bool
	Tags              []string

	Period       Period
	Nugget       Nugget
	HTeam        TeamScoreboard
	VTeam        TeamScoreboard
	Video        Video
	Broadcasters []Broadcaster
}

func GamesByDay(day time.Time) ([]Game, error) {
	year := strconv.Itoa(day.Year())
	all_games, err := GamesByYear(year)

	if err != nil {
		return []Game{}, err
	}

	games := []Game{}
	for _, game := range all_games {
		if CompareDay(game.StartDateEastern, day) {
			games = append(games, game)
		}
	}
	return games, nil
}

func GamesByYear(year string) ([]Game, error) {
	all_games := []Game{}
	raw_json, err := MakeRequest(fmt.Sprintf("/prod/v2/%s/schedule.json", year))
	if err != nil {
		return all_games, err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(raw_json), &result)

	league_schedule := unwrapPath(result, []string{"league"})
	games_raw := unwrapMap(league_schedule, func(league_name string, d interface{}) interface{} {
		games := unwrapArray(d, func(i int, game_raw interface{}) interface{} {

			game := Game{}
			mapstructure.WeakDecode(game_raw, &game)
			game.League = league_name
			game.Video = LoadVideoFromSchedule(game_raw)
			game.Broadcasters = LoadBroadcastersFromSchedule(game_raw)

			// Convert the Eastern time to date structure
			layout := "20060102"
			eastern_time := game_raw.(map[string]interface{})["startDateEastern"]
			game.StartDateEastern, err = time.Parse(layout, eastern_time.(string))
			return game
		})
		return games
	})

	for _, games_array := range games_raw {
		for _, game_raw := range games_array.([]interface{}) {
			game := game_raw.(Game)
			all_games = append(all_games, game)
		}
	}
	return all_games, nil
}
