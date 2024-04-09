// repeatPattern.go
package models

type RepeatPattern string

const (
	None      RepeatPattern = "none"
	Daily     RepeatPattern = "daily"
	Weekday   RepeatPattern = "weekday"
	Weekend   RepeatPattern = "weekend"
	WeeklyMon RepeatPattern = "weekly_mon"
	WeeklyTue RepeatPattern = "weekly_tue"
	WeeklyWed RepeatPattern = "weekly_wed"
	WeeklyThu RepeatPattern = "weekly_thu"
	WeeklyFri RepeatPattern = "weekly_fri"
	WeeklySat RepeatPattern = "weekly_sat"
	WeeklySun RepeatPattern = "weekly_sun"
	Monthly   RepeatPattern = "monthly"
	Yearly    RepeatPattern = "yearly"
	Custom    RepeatPattern = "custom"
)
