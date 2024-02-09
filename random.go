package main

import (
	"fmt"
	"time"
)

func main() {
	LayoutDateTimeMakeFormat := "2006-01-02T15:04:05.000Z"
	LayoutDateTimeFormat := `2006-01-02 15:04:05`
	createdAt, _ := time.Parse(LayoutDateTimeMakeFormat, "2025-10-23T17:38:05.000Z")
	createdAt = createdAt.Add(time.Second * time.Duration(2629746))
	fmt.Println(createdAt.Format(LayoutDateTimeFormat))
}
