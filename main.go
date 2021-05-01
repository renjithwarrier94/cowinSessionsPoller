package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	requestURL          = "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByDistrict"
	bangaloreDistrictID = "265"
)

func main() {
	// HTTP Client
	client := http.Client{}
	// Create a context
	ctx, cancel := context.WithCancel(context.Background())
	// Create a wait group
	wg := sync.WaitGroup{}
	// Slots for today
	wg.Add(1)
	go func() {
		defer wg.Done()
		queryRunningSessionsContinuously(ctx, &client, 0*time.Hour, "today")
	}()
	// Slots for tomorrow
	wg.Add(1)
	go func() {
		defer wg.Done()
		queryRunningSessionsContinuously(ctx, &client, 24*time.Hour, "tomorrow")
	}()
	// Slots for day after tomorrow
	wg.Add(1)
	go func() {
		defer wg.Done()
		queryRunningSessionsContinuously(ctx, &client, 48*time.Hour, "dayAfterTomorrow")
	}()
	// Slots for 2 days from now
	wg.Add(1)
	go func() {
		defer wg.Done()
		queryRunningSessionsContinuously(ctx, &client, 72*time.Hour, "twoDaysFromNow")
	}()
	// Wait until Ctrl+C
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	// Wait until it happens
	<-c
	// Cancel the conntext
	cancel()
	// Wait for it to end
	wg.Wait()
}

func queryRunningSessionsContinuously(ctx context.Context, client *http.Client, requestFor time.Duration, timeFrameName string) {
	// Run continuously
	for {
		// Get current time
		now := time.Now()
		// Get te required time
		requestDate := now.Add(requestFor)
		_, err := getEligileOpenSlotsFor(client, requestDate)
		if err != nil {
			log.Print(err)
		}
		// Log the run
		log.Printf("Executed for %s", timeFrameName)
		// Wait for interval or until the context is done
		select {
		case <-ctx.Done():
			return
		case <-time.After(6 * time.Second):
		}
	}
}

func getEligileOpenSlotsFor(client *http.Client, requestDate time.Time) ([]SessionModel, error) {
	// Fire for the given date
	sessions, err := fireRequest(
		&http.Client{},
		fmt.Sprintf("%02d-%02d-%d\n", requestDate.Day(), requestDate.Month(), requestDate.Year()),
	)
	if err != nil {
		return nil, err
	}

	eligibleSessions := []SessionModel{}
	// Look for sessions with min age <= 27
	for _, session := range sessions {
		if session.MinAgeLimit <= 27 {
			eligibleSessions = append(eligibleSessions, session)
		}
	}

	return eligibleSessions, nil
}

func fireRequest(client *http.Client, date string) ([]SessionModel, error) {
	request, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	// Add query params
	q := request.URL.Query()
	q.Add("district_id", bangaloreDistrictID)
	q.Add("date", date)
	request.URL.RawQuery = q.Encode()
	// Add required headers
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Accept-Language", "en_US")
	// Fire the request
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Read resp body
	respBody, _ := ioutil.ReadAll(resp.Body)
	// Deserialise the input
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("non 200 response code %d", resp.StatusCode)
	}

	var sessions SessionsResponse

	err = json.Unmarshal(respBody, &sessions)
	if err != nil {
		return nil, err
	}

	return sessions.Sessions, nil
}
