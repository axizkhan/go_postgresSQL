package user

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T){

	dob := time.Date(
		2000,
		time.January,
		1,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	age := calculateAge(dob)

	if age <= 0 {
		t.Errorf("expected valid age, got %d", age)
	}
}