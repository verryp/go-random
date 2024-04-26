package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	Station   string   `json:"station"`
	Schedules []Detail `json:"schedules"`
}

type Detail struct {
	TrainNumber   string `json:"train_number"`
	DepartureTime string `json:"departure_time"`
	ArrivalTime   string `json:"arrival_time"`
	Destination   string `json:"destination"`
}

func main() {
	response := []byte(`{
  "status": "success",
  "data": {
    "station": "Grand Central Station",
    "schedules": [
      {
        "train_number": "A123",
        "departure_time": "2024-02-22T11:00:00",
        "arrival_time": "2024-02-22T13:30:00",
        "destination": "Cityville"
      },
      {
        "train_number": "C789",
        "departure_time": "2024-02-22T18:45:00",
        "arrival_time": "2024-02-22T21:15:00",
        "destination": "Village Town"
      }
    ]
  }
}`)

	temp := Response{}

	//b, err := json.Marshal(response)
	//fmt.Println("err", err)
	err := json.Unmarshal(response, &temp)
	fmt.Println("err", err)

	fmt.Println("ddd", findTrainSchedule(temp, "Cityville"))

	//fmt.Println("ddd", temp)
}

func findTrainSchedule(destination Response, dest string) string {
	for _, dd := range destination.Data.Schedules {
		if dd.Destination == dest {
			return fmt.Sprintf("%v", dd)
		}
	}

	return ""
}
