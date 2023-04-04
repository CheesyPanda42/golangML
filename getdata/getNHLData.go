// This module is used to get the NHL data from the NHL API
// NHL stats API documentation is at https://gitlab.com/dword4/nhlapi
// NHL stats API base url is https://statsapi.web.nhl.com/api/v1

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get data from the specified URL
func getData(url string) []byte {
	// Get the data from the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// Read the data
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}

// Get the NHL data.
// Takes a string "year" as a parameter.
// Returns the game results of all games from the specified year to present day
func getNHLData(year string) {
	// Get the data from the NHL API
	data := getData("https://statsapi.web.nhl.com/api/v1/schedule?startDate=" + year + "-01-01&endDate=" + year + "-12-31")

	// Unmarshal the data into a struct
	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	// Get the dates from the data
	dates := result["dates"].([]interface{})

	// Loop through the dates
	for _, date := range dates {
		// Get the games from the date
		games := date.(map[string]interface{})["games"].([]interface{})

		// Loop through the games
		for _, game := range games {
			// Get the game data
			gameData := game.(map[string]interface{})

			// Get the game ID
			gameID := gameData["gamePk"].(float64)

			// Get the game date
			gameDate := gameData["gameDate"].(string)

			// Get the game status
			gameStatus := gameData["status"].(map[string]interface{})["abstractGameState"].(string)

			// Get the game type
			gameType := gameData["gameType"].(string)

			// Get the game season
			gameSeason := gameData["season"].(string)

			// Get the game teams
			gameTeams := gameData["teams"].(map[string]interface{})

			// Get the home team
			homeTeam := gameTeams["home"].(map[string]interface{})["team"].(map[string]interface{})["name"].(string)

			// Get the away team
			awayTeam := gameTeams["away"].(map[string]interface{})["team"].(map[string]interface{})["name"].(string)

			// Get the game score
			gameScore := gameData["teams"].(map[string]interface{})["home"].(map[string]interface{})["score"].(float64)

			// Print the game data
			fmt.Println(gameID, gameDate, gameStatus, gameType, gameSeason, homeTeam, awayTeam, gameScore)
		}
	}
}

// Main function
func main() {
	// Get the NHL data
	getNHLData("1999")
}
