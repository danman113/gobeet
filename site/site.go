package site

import (
	"time"
)

type Website struct {
	Url      string `json: url`
	Interval int    `json: interval`
	Pages    []Page `json: pages`
}

type Page struct {
	Url       string `json: url`
	Status    int    `json: status`
	Method    string `json: method`
	Timeout   int    `json: timeout`
	Duration  int    `json: duration`
	DownSince *time.Time
}
