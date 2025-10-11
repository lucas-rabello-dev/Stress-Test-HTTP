package model

import (
	"time"

	"github.com/lucas-rabello-dev/Stress-Test-HTTP/internal/model"
)

// entidades principais

type DataFlag struct {
	URL string
	Requests int
	Time time.Duration
	JsonFileName string
	Method string
}

func addData(url string, requests int, time time.Duration, jsonFileName string, mothod string) *model.DataFlag {
	return &model.DataFlag{
		
	}
}