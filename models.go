package main

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
