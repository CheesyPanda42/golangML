// This module is used to get the NHL data from the NHL API
// NHL stats API documentation is at https://gitlab.com/dword4/nhlapi
// NHL stats API base url is https://statsapi.web.nhl.com/api/v1

package main

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// Create wait group to wait for all goroutines to finish
var wg sync.WaitGroup

func getData(year int) {
	// Defer the wait group
	defer wg.Done()

	// Set the start date to January 1st of the current year
	startDate := fmt.Sprintf("%d-01-01", year)

	// Set the end date to December 31 of the current year
	endDate := fmt.Sprintf("%d-12-31", year)
	// Print the start and end dates
	fmt.Println(startDate, endDate)

	// Print a statement to show that the data is being retrieved
	fmt.Println("Retrieving data for year", year)

	// Make string witi the url
	url := "https://statsapi.web.nhl.com/api/v1/schedule?startDate=" + startDate + "&endDate=" + endDate

	// Make http request to the NHL API
	resp, err := http.Get(url)

	// Check for errors
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Convert the response body to a byte slice
	body, err := ioutil.ReadAll(resp.Body)

	// Check for errors
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Write the JSON response to "nhlData_i.json", where i is the year being looped through
	err = ioutil.WriteFile("data/nhlData_"+fmt.Sprintf("%d", year)+".json", body, 0644)

	// Print a statement to show that the data was written to the file
	fmt.Println("Data written to file for year", year)

	// Check for errors
	if err != nil {
		fmt.Println("Error:", err)
	}

}

func main() {
	// Get NHL game data from 2000-01-01 to present day, from the NHL API at https://statsapi.web.nhl.com/api/v1/schedule?

	// Check for the correct number of arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run getNHLData.go <start year> <end year>")
		os.Exit(1)
	}

	// Convert start date and end date to integers
	startDate, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
	}

	endDate, err := strconv.Atoi(os.Args[2])

	// Check for errors
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Loop throug the years from 1990 to 2020
	for i := startDate; i < endDate; i++ {
		// Print the year
		fmt.Println(i)
		// Get the data from the NHL API as a goroutine and add it to the wait group
		go getData(i)
		wg.Add(1)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
