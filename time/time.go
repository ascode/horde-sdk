package time

import "time"

// GetFirstDateOfTheWeek 获取本周周一的日期
func GetFirstDateOfTheWeek(t time.Time) time.Time {

	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).
		AddDate(0, 0, offset)

}

// GetFirstDayOfNextWeek 获取下周周一
func GetFirstDayOfNextWeek(t time.Time) time.Time {

	return GetFirstDateOfTheWeek(t).
		AddDate(0, 0, 7)

}

// GetEndDateOfTheWeek 获取本周周日
func GetEndDateOfTheWeek(t time.Time) time.Time {

	return GetFirstDateOfTheWeek(t).
		AddDate(0, 0, 6)

}

// GetFirstDayOfLastWeek 获取上周的周一日期
func GetFirstDayOfLastWeek(t time.Time) time.Time {
	thisWeekMonday := GetFirstDateOfTheWeek(t)

	return thisWeekMonday.AddDate(0, 0, -7)
}

// GetEndDayOfNextWeek 获取下周周日
func GetEndDayOfNextWeek(t time.Time) time.Time {
	return GetEndDateOfTheWeek(t).AddDate(0, 0, -7)
}
