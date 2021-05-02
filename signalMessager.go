package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type SignalMessager struct {
	clientCode   string
	clientNumber string
	httpClient   *http.Client
}

func GetSignalMessagerClients(httpClient *http.Client) []*SignalMessager {
	var messengerClients []*SignalMessager
	for i, phno := range phoneNumbers {
		client := newSignalMessager(apiKeys[i], phno, httpClient)
		messengerClients = append(messengerClients, client)
	}

	return messengerClients
}

func newSignalMessager(clientCode, clientNumber string, client *http.Client) *SignalMessager {
	return &SignalMessager{
		clientCode:   clientCode,
		clientNumber: clientNumber,
		httpClient:   client,
	}
}

func (s *SignalMessager) SendSessions(sessions Sessions, date time.Time) {
	// Construct the request
	request, err := http.NewRequest("GET", callMeBotURL, nil)
	if err != nil {
		log.Printf("Error while pushing to signal: %v", err)
		return
	}
	// Add query params
	q := request.URL.Query()
	q.Add("phone", s.clientNumber)
	q.Add("apikey", s.clientCode)
	// Loop through them and send it
	q.Add(
		"text",
		fmt.Sprintf(
			"Slots for %s: \n%s",
			// Add the current time in dd-mm-yyyy format
			fmt.Sprintf("%02d-%02d-%d\n", date.Day(), date.Month(), date.Year()),
			// Convert to string
			sessions.String(),
		),
	)
	// Add it
	request.URL.RawQuery = q.Encode()
	// Send the request
	resp, err := s.httpClient.Do(request)
	if err != nil {
		log.Printf("Error while pushing to signal: %v", err)
		return
	}
	// Close the body
	defer resp.Body.Close()
	// Check the status code
	if resp.StatusCode != 200 {
		// Print the result
		log.Printf("Non 200 response when trying to push to signal: %d", resp.StatusCode)
		log.Printf("Failed to send: \n%s", sessions.String())
	}
}
