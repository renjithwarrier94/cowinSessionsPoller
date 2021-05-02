# Introduction

This is a simple component that is written with the intention of polling the list of availale vaccine slots in a district and send out a notification when the sessions matching certain criteria is found.

# Configuration

configs.go file can be used to configure the working of this component. The district to poll for, the filter criteria (right now it checks for minimum age and the vaccine used). The phone numbers and its corresponding API keys are also configured here.

# Working

This component polls the public sessions [API](https://apisetu.gov.in/public/api/cowin) exposed by cowin.