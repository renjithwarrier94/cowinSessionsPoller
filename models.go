package main

import (
	"fmt"
	"strings"
)

type SessionModel struct {
	CenterId          int      `json:"center_id"`
	Name              string   `json:"name"`
	StateName         string   `json:"state_name"`
	DistrictName      string   `json:"district_name"`
	BlockName         string   `json:"block_name"`
	Pincode           int      `json:"pincode"`
	From              string   `json:"from"`
	To                string   `json:"to"`
	Lat               int      `json:"lat"`
	Long              int      `json:"long"`
	FeeType           string   `json:"fee_type"`
	SessionID         string   `json:"session_id"`
	Date              string   `json:"date"`
	AvailableCapacity int      `json:"availabble_capacity"`
	Fee               string   `json:"fee"`
	MinAgeLimit       int      `json:"min_age_limit"`
	Vaccine           string   `json:"vaccine"`
	Slots             []string `json:"slots"`
}

type SessionsResponse struct {
	Sessions []SessionModel `json:"sessions"`
}

// Define an alias for a slice of SessionModel
type Sessions []SessionModel

func (sessions Sessions) String() string {
	builder := strings.Builder{}

	for _, s := range sessions {
		builder.WriteString(fmt.Sprintf(
			"\nCenter name: %s\nPincode: %d\nFee: %s\nVaccine: %s\nSlots: %v\n",
			s.Name,
			s.Pincode,
			s.Fee,
			s.Vaccine,
			s.Slots,
		),
		)
	}

	return builder.String()
}
