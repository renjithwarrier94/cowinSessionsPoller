# Introduction

This is a simple component that is written with the intention of polling the list of availale vaccine slots in a district and send out a notification when the sessions matching certain criteria is found.

# Configuration

configs.go file can be used to configure the working of this component. The district to poll for, the filter criteria (right now it checks for minimum age and the vaccine used). The phone numbers and its corresponding API keys are also configured here.

# Working

This component polls the public sessions [API](https://apisetu.gov.in/public/api/cowin) exposed by cowin. It sets up workers to poll the sessions avaible today, tommorrow, day after and 2 days from now. If it finds sessions that match our filter criteria, it initiates a notification via Signal App. 

# Running

The simplest way to run is to just install go, clone the repo, change to the project directory and run using the command:
```
go run .
```

If you want to deploy this somewhere, the project can be built using:
```
go build .
```
Then the built binary can be run using the command:
```
./cowinSessionsPoller
```