# Sylabs Cloud Engineering Homework

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


## Testing

- Unit Testing - created 
 
- Test Coverage -  is the percentage of your code covered by test suit. In layman’s language, it is the measurement of how many lines of code in your package were executed when you ran your test suit (compared to total lines in your code). Go provide built-in functionality to check your code coverage.


- Using Assertions - [Testify](https://github.com/stretchr/testify) Go does not provide any built-in package for assertions  so i decided to use 3rd party  

