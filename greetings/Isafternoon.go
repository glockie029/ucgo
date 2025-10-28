package greetings

import "time"

func Afternoon() bool {
	currenttime := time.Now().Hour()
	if currenttime > 12 && currenttime < 18 {
		return true
	}
	return false
}
