package main

import (
	"fmt"
	"time"
)

func main() {

	//now := time.Now()
	//weekday := now.Weekday()
	//fmt.Println(weekday, int(weekday))
	//fmt.Println(int(time.Monday), time.Monday)
	//
	//year, month, _ := now.Date()
	//thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	//start := thisMonth.AddDate(0, -1, 0).Format("2006-01-02 15:04:05")
	//end := thisMonth.AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
	//fmt.Println(start, end)

	//fmt.Printf("type:%#v\n", append(make([]int, 0), 1, 2))

	dateList := make(map[string]interface{})
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	fmt.Println(weekday)
	fmt.Println(int(time.Monday))

	dateList["y_begin"] = now.AddDate(0, 0, int(time.Monday)-weekday-14).Format("2006-01-02") + " 00:00:00"
	dateList["y_end"] = now.AddDate(0, 0, int(time.Monday)-weekday-1).Format("2006-01-02") + " 23:59:59"
	dateList["y1_begin"] = now.AddDate(0, 0, int(time.Monday)-weekday-21).Format("2006-01-02") + " 00:00:00"
	dateList["y1_end"] = now.AddDate(0, 0, int(time.Monday)-weekday-8).Format("2006-01-02") + " 23:59:59"

	fmt.Println(dateList)
}
