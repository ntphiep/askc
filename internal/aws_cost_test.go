package internal

import (
	"testing"
)

// Note: Real AWS calls are not easily testable without mocking AWS SDK.
// This is a placeholder for future mock-based tests.
func TestGetDailyCost_Sample(t *testing.T) {
	_, err := GetDailyCost()
	if err != nil {
		t.Logf("Expected error if AWS credentials are missing: %v", err)
	}
}

// func TestGetYesterdayCostDetails(t *testing.T) {
// 	details, err := GetYesterdayCostDetails()
// 	if err != nil {
// 		t.Fatalf("Error getting cost details: %v", err)
// 	}
// 	t.Log(details)
// }
