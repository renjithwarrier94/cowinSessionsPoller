package main

func isSessionEligible(session SessionModel) bool {
	// Check the minimum age
	if session.MinAgeLimit > minAge {
		return false
	}
	// Check the vaccine type
	if session.Vaccine != vaccineType {
		return false
	}
	return true
}
