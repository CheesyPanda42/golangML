// Using GoLearn to train a model to predict NHL game outcomes based on the number of vowels in the team names
// The NHL Stats API base url is https://statsapi.web.nhl.com/api/v1
// The NHL Stats API documentation is at https://gitlab.com/dword4/nhlapi






package main

import (
	"fmt"
	"log"
	// Import Golearn packages
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/linear_models"
)

func main() {
	// Load the csv file into a dataframe
	// CSV file is from Kaggle: https://www.kaggle.com/deepmatrix/imdb-5000-movie-dataset
	// Download CSV file and place it in the same directory as this file
	
	

	// Load the training set
	trainingData, err := base.ParseCSVToInstances("movie_metadata_training.csv", true)