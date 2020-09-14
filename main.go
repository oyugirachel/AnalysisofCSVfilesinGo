package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Statistic represents a row of data
type Statistic struct {
	BuildID             string    `json:"build_id"`
	UserID              string    `json:"user_id"`
	TimeRequestReceived time.Time `json:"time_request_received"`
	TimeRequestBegan    time.Time `json:"time_request_began"`
	TimeRequestFinished time.Time `json:"time_request_finished"`
	BuildDeleted        bool      `json:"build_deleted"`
	BuildProcess        int       `json:"build_process"`
	ImageSize           int       `json:"image_size"`
}

// pointing  to Statistic struct
var stats []*Statistic
//Time request started
const(
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006 "
)

func main() {

	//Reading the CSV
	if err := ReadCsv("stats.csv"); err != nil {
		log.Fatal(err)
	}
	//Inputing the time range to a variable t
	startdate := "2018-10-6"
	enddate := "2018-11-10"
	s, _ := time.Parse(layoutISO, startdate)
	e, _ := time.Parse(layoutISO, enddate)
  //Refers to the Build Execution time

	finished := BuildExecuteTime(s, e)
	fmt.Printf("%d builds finished in provided range\n\n", finished)

	user := UserBuildRemoteService()
	fmt.Println(user)

	build := BuildSuccessRate()
	fmt.Printf("%d builds not succeeding\n", build)

}

func BuildExecuteTime(start, end time.Time) int {
	var count = 0

	for _, stat := range stats {
		if inTimeSpan(start, end, stat.TimeRequestFinished) {
			count++
		}
	}

	return count
}

func UserBuildRemoteService() int {
	var count = 0

	for _, stat := range stats {


		if stat.UserID  == stat.UserID {
			//count++

			fmt.Println(stat.UserID)
		}

	}

	return count
}

func BuildSuccessRate() int {

	var count = 0

	for _, stat := range stats {

		if stat.BuildProcess > 0 {
			count++
			fmt.Printf("%d top exit codes\n", stat.BuildProcess)

		}
	}

	return count
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

// ReadCsv accepts a file and reads an array of Statistics
func ReadCsv(filename string) error {
	// Open CSV file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read File into a Variable To read in the CSV data, we create a new CSV reader that reads from the input file.
	//The CSV reader is aware of the CSV data format. It separates the input stream into rows and columns, and returns a slice of slices of strings.
	lines, err := csv.NewReader(file).ReadAll() //Read the whole file at once. (We don’t expect large files.)
	if err != nil {
		return err
	}  //We check for any error
    //Now that the data is read in, we can loop over data, and read from or write to each line slice as needed.
	for _, line := range lines {
		received, _ := time.Parse(time.RFC3339, line[2])
		began, _ := time.Parse(time.RFC3339, line[3])
		finished, _ := time.Parse(time.RFC3339, line[4])
		deleted, _ := strconv.ParseBool(line[5])
		process, _ := strconv.Atoi(line[6])
		size, _ := strconv.Atoi(line[7])

		stat := &Statistic{
			BuildID:             line[0],
			UserID:              line[1],
			TimeRequestReceived: received,
			TimeRequestBegan:    began,
			TimeRequestFinished: finished,
			BuildDeleted:        deleted,
			BuildProcess:        process,
			ImageSize:           size,
		}

		stats = append(stats, stat) //appending the data to the stats file
	}

	return nil
}
