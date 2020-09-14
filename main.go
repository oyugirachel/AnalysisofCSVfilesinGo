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

func main() {

	//Reading the CSV
	if err := ReadCsv("stats.csv"); err != nil {
		log.Fatal(err)
	}


	finished := BuildExecuteTime(time.Now(), time.Now())
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

	// Read File into a Variable
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}

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

		stats = append(stats, stat)
	}

	return nil
}
