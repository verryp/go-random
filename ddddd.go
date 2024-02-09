package main

import (
	"fmt"
	"time"
)

func endDate(t time.Time, wd time.Weekday) time.Time {
	next := int((wd - t.Weekday() + 7) % 7)
	y, m, d := t.Date()
	return time.Date(y, m, d+next+1, 0, 0, 0, -1, t.Location())
}

func main() {
	now := time.Now().Round(0)
	fmt.Println(now, now.Weekday())

	end := endDate(now, time.Sunday)
	fmt.Println(end, end.Weekday())

	ddd := time.Date(2023, 12, 7, 0, 0, 0, 0, time.Local)
	fmt.Println("ddd", ddd)

	//tn := time.Now()
	//fmt.Println("nowww", tn)
	//ddd := time.Second - time.Duration(1)
	fmt.Println("modify", ddd.Add(-(time.Second*1)+time.Hour*24))
}
