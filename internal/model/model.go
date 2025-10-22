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
}

func AddData(url *string, requests *int, time *time.Duration, jsonFileName *string) *DataFlag {
	return &DataFlag{
		URL: *url,
		Requests: *requests,
		Time: *time,
		JsonFileName: *jsonFileName,
	}
}