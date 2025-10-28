package main

import (
	"fmt"
	"time"
)

var week time.Duration

func FunctionFromFile1() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	t = time.Now().UTC()
	fmt.Println(t)
	week = 60 * 60 * 24 * 7 * 1e9
	fmt.Println(week)
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now)
}
