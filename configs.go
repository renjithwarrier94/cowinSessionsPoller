package main

// Configs related to polling the open sessions API
// Inorder to find the district id, use the APIs at:
// https://apisetu.gov.in/public/api/cowin
const (
	// The URL to get sessions from
	requestURL = "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByDistrict"
	// District ID for BBMP
	bangaloreDistrictID = "294"
	// Run interval
	runIntervalSeconds = 60
	// Sleep interval when a match is found
	sleepIntervalOnSuccessSeconds = 10 * 60
)

// Configs related to the messenger API
// Inorder to set it up, follow the guide at:
// https://www.linkedin.com/posts/abhinav-dinesh-60b770a4_want-to-get-notified-when-theres-a-covid-activity-6793910597353402368-ojox
const (
	callMeBotURL = "https://api.callmebot.com/signal/send.php"
)

// The phone numbers and apikeys to be used in the call me ot API
var (
	// List of phone numbers to notify
	phoneNumbers = [...]string{
		"+919876543210",
	}
	// Corresponding Api Keys
	apiKeys = [...]string{
		"123456",
	}
)

// Configs related to filtering open sessions to identify eligible sessions
const (
	// Set the vaccine type
	vaccineType = "COVISHIELD"
	// Vaccine age
	minAge = 26
)
