package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

/*
 */

const (
	requestURL          = "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByDistrict"
	bangaloreDistrictID = "265"
)

func main() {
	fireRequest(&http.Client{}, "04-05-2021")
}

func fireRequest(client *http.Client, date string) {
	request, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Print(err)
		return
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
		log.Printf("Error while firing the query: %v", err)
	}
	defer resp.Body.Close()
	// Read resp body
	respBody, _ := ioutil.ReadAll(resp.Body)
	// Deserialise the input
	if resp.StatusCode != 200 {
		log.Printf("Non success status code: %d", resp.StatusCode)
		return
	}

	var sessions SessionsResponse

	err = json.Unmarshal(respBody, &sessions)
	if err != nil {
		log.Printf("Invalid response body: %s", respBody)
		log.Printf("Error: %v", err)
		return
	}

	if len(sessions.Sessions) == 0 {
		log.Print("No available sessions")
		return
	}

	log.Printf("Slots: %v", sessions.Sessions[0].Slots)
}
