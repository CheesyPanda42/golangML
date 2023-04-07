// This is a Go application that parses the NHL data gathered by getNHLData
// The data files provided are JSON output from the NHL API
// For each game in the data, we want to create a feature vector that contains the number of vowels in the home team name and the number of vowels in the away team name, and a label that is 1 if the home team won and 0 if the home team lost
// The feature vector and label will be written to a CSV file, which will be used to train a model in ml/main.go

// Import packages
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
)

// Define a struct to hold the data from the NHL API
// We are only interested in the gameDate, the home team name and score, and the away team name and score

type NHLData struct {
	Dates []struct {
		Games []struct {
			GameDate string `json:"gameDate"`
			Teams    struct {
				Away struct {
					Score int `json:"score"`
					Team  struct {
						Name string `json:"name"`
					} `json:"team"`
				} `json:"away"`
				Home struct {
					Score int `json:"score"`
					Team  struct {
						Name string `json:"name"`
					} `json:"team"`
				} `json:"home"`
			} `json:"teams"`
		} `json:"games"`
	} `json:"dates"`
}

// Utility function for error handing
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Create wait group to wait for all goroutines to finish
var wg sync.WaitGroup

func parseData(year int) {
	// Print a statement to show that the data is being parsed
	fmt.Println("Parsing data for year", year)

	defer wg.Done()
	// Read in the data from the file
	dataFilePath := "../getdata/data/nhlData_" // Path to the data files

	dataFile, err := os.Open(dataFilePath + fmt.Sprintf("%d", year) + ".json")
	check(err)

	// Create a variable of type NHLData to hold the data
	var nhlData NHLData

	// Decode the JSON data into the NHLData struct
	err = json.NewDecoder(dataFile).Decode(&nhlData)
	check(err)

	// Create a CSV file to write the data to
	csvFile, err := os.Create("../getdata/data/nhlData_" + fmt.Sprintf("%d", year) + ".csv")
	check(err)

	// Write the header row to the CSV file
	_, err = csvFile.WriteString("homeTeamName,homeTeamScore,awayTeamName,awayTeamScore,homeTeamVowels,awayTeamVowels,homeTeamWon\n")
	check(err)

	// For each date in the data, loop through the games and create a feature vector and label for each game
	for _, date := range nhlData.Dates {
		for _, game := range date.Games {
			// Get the home team name and score
			homeTeamName := game.Teams.Home.Team.Name
			homeTeamScore := game.Teams.Home.Score

			// Get the away team name and score
			awayTeamName := game.Teams.Away.Team.Name
			awayTeamScore := game.Teams.Away.Score

			// Create a feature vector that contains the number of vowels in the home team name and the number of vowels in the away team name
			homeTeamVowels := 0
			awayTeamVowels := 0
			for _, char := range homeTeamName {
				if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
					homeTeamVowels++
				}
			}
			for _, char := range awayTeamName {
				if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
					awayTeamVowels++
				}
			}

			// Create a label that is 1 if the home team won and 0 if the home team lost
			homeTeamWon := 0
			if homeTeamScore > awayTeamScore {
				homeTeamWon = 1
			}
			// Print a statement to show that the data is being written to the CSV file
			fmt.Println("Writing data for", homeTeamName, "vs", awayTeamName, "to CSV file")

			// Write the feature vector and label to the CSV file
			_, err = csvFile.WriteString(fmt.Sprintf("%s,%d,%s,%d,%d,%d,%d\n", homeTeamName, homeTeamScore, awayTeamName, awayTeamScore, homeTeamVowels, awayTeamVowels, homeTeamWon))
			check(err)
		}
	}
}

// Read in data and write to CSV file
func main() {

	// Get command line arguments for what years to parse. The first argument is the start year, the second argument is the end year. Convert to integers
	// Check for the correct number of arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run parseNHLdata.go <start year> <end year>")
		os.Exit(1)
	}

	startYear, err := strconv.Atoi(os.Args[1])
	check(err)
	endYear, err := strconv.Atoi(os.Args[2])
	check(err)

	// Loop through the years
	for i := startYear; i < endYear; i++ {
		wg.Add(1)
		// Create a go routine to read in the data and write it to a CSV file
		go parseData(i)
	}
	wg.Wait()
}
