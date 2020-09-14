# Sylabs Cloud Engineering Homework

## Overview

This task is not meant to take a substantial amount of time. If you spend more than an hour on it then you are probably taking it further than we intended. Please try to limit yourself to just doing your best within a one hour window. We understand that some people naturally will want to spend more time to make the solution more elegant or to over-deliver on features or deliverables, but trust us that it is not necessary. We are not looking for one correct solution, but rather are trying to learn about how you think, how you work, and more basically if you can deliver a solution that accomplishes the task.

Your solution to the task is not work that will be used after the interview process. This task is only hypothetical, and while it is related to work you could be asked to do if you join the team, the scale and details of the task would be substantially different than this much simpler task we have formulated below.

For more information work sample tests and why we use them please see Work Sample Tests.

## Description

Our marketing department would like some usage information regarding our Sylabs Cloud Remote Builder service. We would like you to provide a very basic reporting implementation for the data we have saved from the builds that have been executed.

You will be given a CSV file that contains information from our Remote Builder service that executes Singularity builds in our Sylabs Cloud based on user requests. The CSV file contents will consist of fields in this order:

* Unique identifier for each build
* User ID reference the user that submitted the build request
* Time the build request was received (RFC 3339 formatted string)
* Time the build execution began (RFC 3339 formatted string)
* Time the build execution finished (RFC 3339 formatted string)
* Indicator for if the build has been deleted
* Exit code from the build process, >0 indicates failure
* Size of the resulting built image file

The marketing department wants to know how many builds were executed in a time window. For example, how many builds were executed in the last 15 minutes, or in the last day, or between January 1 and January 31, 2018.

The marketing department wants to know which users are using the remote build service the most. Who are the top 5 users and how many builds have they executed in the time window?

The marketing department would like to know the build success rate, and for builds that are not succeeding what are the top exit codes.

## Task

Design a simple solution to the situation described above using the Go programming language and implemented as command-line program. The input will be a CSV file. The output should be to stdout/stderr. Your submission should include your source code, tests, and a README for your solution. In your README please describe your object model, information about the software design choices you made, and your testing approach. Please also include a paragraph on how you would extend your solution to make it better if given more time.

This submission is due no later than 24 hours before your scheduled interview time. Please send a tar.gz file containing your code, tests, test fixtures, dependencies and your README to your Sylabs HR contact via email.
