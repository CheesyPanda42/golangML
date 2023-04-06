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
func getNHLData(year string) map[string]interface{} {
	// Get the data from the NHL API
	data := getData("https://statsapi.web.nhl.com/api/v1/schedule?startDate=" + year + "-01-01&endDate=" + year + "-12-31")

	// Unmarshal the data into a struct
	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	// Return the data
	return result
}

// Main function
func main() {
	// Get the NHL data

}
