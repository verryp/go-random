// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

const (
	// FormatDate ...
	FormatDate = `2006-01-02`
	// TimeAsiaJakarta ...
	TimeAsiaJakarta = `Asia/Jakarta`
)

var (
	loc, _ = time.LoadLocation(TimeAsiaJakarta)
)

// ParseUTC7 ...
func ParseUTC7(timeFormat string, value string) (time.Time, error) {
	timeUTC7, err := time.ParseInLocation(timeFormat, value, loc)
	if err != nil {
		return time.Now(), err
	}

	return timeUTC7, nil
}

func main() {

	endDate, _ := ParseUTC7(FormatDate, "2024-06-10")
	fmt.Println("before : ", endDate)

	fmt.Println("after : ", endDate.Add(time.Hour*time.Duration(23)+time.Minute*time.Duration(59)+time.Second*time.Duration(59)))
	dd := endDate.Add(-(time.Second * 1) + time.Hour*24)
	fmt.Println("after : ", dd)

	fmt.Println(DateAfter(endDate, endDate))
	fmt.Println(DateAfter(dd, endDate))

	fmt.Println(DateBeforeOrEqual(dd, dd))

}

// DateAfter ...
func DateAfter(start time.Time, end time.Time) bool {
	start, _ = ParseUTC7(FormatDate, start.Format(FormatDate))
	end, _ = ParseUTC7(FormatDate, end.Format(FormatDate))

	return start.After(end)
}

// DateBeforeOrEqual ...
func DateBeforeOrEqual(t1, t2 time.Time) bool {
	t1, _ = ParseUTC7(FormatDate, t1.Format(FormatDate))
	t2, _ = ParseUTC7(FormatDate, t2.Format(FormatDate))

	return t1.Before(t2) || t1.Equal(t2)
}