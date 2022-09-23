package main

import (
	"fmt"
	"time"
)

type Book struct {
	Title      string
	Author     string
	Publisher  string
	ReleasedAt time.Time
	ISBN       string
}

func main() {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	book := Book{
		Title:      "Sample",
		Author:     "Sampleさん",
		Publisher:  "Sample出版",
		ISBN:       "123456789",
		ReleasedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	fmt.Println(book)
}
