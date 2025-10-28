package greetings

import (
	"time"
)

func IsAM() bool {
	currenttime := time.Now()
	currentHour := currenttime.Hour()
	if currentHour < 12 {
		return true
	}
	return false
}
