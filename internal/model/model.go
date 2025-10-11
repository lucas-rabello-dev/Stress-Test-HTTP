package model

import (
	"time"
)

// entidades principais

type DataFlag struct {
	URL string
	Requests int
	Time time.Duration
	JsonFileName string
	Method string
}

func AddData(url *string, requests *int, time *time.Duration, jsonFileName *string, method *string) *DataFlag {
	return &DataFlag{
		URL: *url,
		Requests: *requests,
		Time: *time,
		JsonFileName: *jsonFileName,
		Method: *method,
	}
}