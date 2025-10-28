package greetings

import "time"

func Evening() bool {
	currenttime := time.Now().Hour()
	if currenttime > 18 && currenttime < 24 {
		return true
	}
	return false
}
