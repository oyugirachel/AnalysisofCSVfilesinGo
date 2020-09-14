# Sylabs Cloud Engineering Homework
##Problem Statement
The marketing department wants to know how many builds were executed in a time window. For example, how many builds were executed in the last 15 minutes, or in the last day, or between January 1 and January 31, 2018.

The marketing department wants to know which users are using the remote build service the most. Who are the top 5 users and how many builds have they executed in the time window?

The marketing department would like to know the build success rate, and for builds that are not succeeding what are the top exit codes
## Object Model

```go
package main

import "time"

type Statistic struct {
	BuildId             string    `json:"build_id"`
	UserId              string    `json:"user_id"`
	TimeRequestReceived time.Time `json:"time_request_received"`
	TimeRequestBegan    time.Time `json:"time_request_began"`
	TimeRequestFinished time.Time `json:"time_request_finished"`
	BuildDeleted        bool      `json:"build_deleted"`
	BuildProcess        int       `json:"build_process"`
	ImageSize           int       `json:"image_size"`
}
```


## Software Design
Imports and main
We only use packages from the standard library here(import)
In main(), we sketch out our program flow:
Read the CSV file,
calculate the desired numbers, and output the desired results from the analysis.




##Looping Through the Solution
But before starting to code, our first step is to export the spreadsheet data to CSV. To make things a bit less complicated, we export the data with a semicolon as the column separator.
 ReadCsv(Comma Seperated Values) accepts a file and returns its content as a multi-dimentional type
with lines and each column. Only parses to string type.
 The CSV file is opened using the os.open function to read the data
 This whole process involves looping through the whole data of the CSV file
 The marketing builds  data would best be loaded into the above Statistic Struct. Each property of the Statistic  structure have a JSON annotation which will allow for easy printing to JSON after it had been loaded.
 I had to use RFC(https://tools.ietf.org/html/rfc3339') to work out the time stamps to convert them to normal ones for they are in the Unix format
  I then used map to map the time and used a for loop to loop through the data to find out the range and builds taken within a specified time.
 of TimeRequestReceived time.Time `json:"time_request_received"`
    	TimeRequestBegan    time.Time `json:"time_request_began"`
    	TimeRequestFinished time.Time `json:"time_request_finished"`
    So I have to get the difference of : TimeRequestFinished - TimeRequestBegan to get the overall time.
 ##End Point
 The whole function is  counting  number of builds not succeeding  and their exit Code
 ##Dependencies
 Go Standard Library
 '''
 "encoding/csv"
 	"fmt"
 	"log"
 	"os"
 	"strconv"
 	"time"
 
 '''
  
## Testing

- Unit Testing - created 
 
- Test Coverage -  is the percentage of your code covered by test suit. In layman’s language, it is the measurement of how many lines of code in your package were executed when you ran your test suit (compared to total lines in your code). Go provide built-in functionality to check your code coverage.


- Using Assertions - [Testify](https://github.com/stretchr/testify) Go does not provide any built-in package for assertions  so i decided to use 3rd party  
#Conclusion
This project could be improved more if analysed carefully for example for it to be able to output the number of ids repeated instead of displaying them concurrently and in knowing the real users using the remote build the most.
According to the stats.csv We are getting all the ids being repeated
When you just run the code as it is You get 0 builds succeeded but when you go and toggle with the builds in the csv for example you try toggling to 1 or 2 it outputs 1 build and 2 builds not succeeded respectively.
In the time function(func inTimeSpan) you just pass the time and it will output the data
This whole projects involves extraction and analysis of the data.
