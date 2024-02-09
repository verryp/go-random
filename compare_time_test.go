package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCalculateExpiredAt(t *testing.T) {
	t.Parallel()

	expirationDays := 3
	testCases := []struct {
		input time.Time
		start time.Time
		end   time.Time
		want  string
	}{
		{
			input: time.Now(),
			start: time.Date(2023, 07, 24, 0, 0, 0, 0, time.Local),
			end:   time.Date(2023, 07, 24, 0, 0, 0, 0, time.Local),
			want:  "2023-07-27",
		},
		{
			input: time.Now(),
			start: time.Date(2023, 07, 24, 0, 0, 0, 0, time.Local),
			end:   time.Date(2023, 07, 25, 0, 0, 0, 0, time.Local),
			want:  "2023-07-28",
		},
		{
			input: time.Now(),
			start: time.Date(2023, 07, 24, 0, 0, 0, 0, time.Local),
			end:   time.Date(2023, 07, 26, 0, 0, 0, 0, time.Local),
			want:  "2023-07-31",
		},
		{
			input: time.Now(),
			start: time.Date(2023, 07, 21, 0, 0, 0, 0, time.Local),
			end:   time.Date(2023, 07, 25, 0, 0, 0, 0, time.Local),
			want:  "2023-07-28",
		},
		{
			input: time.Now(),
			start: time.Date(2023, 07, 25, 0, 0, 0, 0, time.Local),
			end:   time.Date(2023, 07, 30, 0, 0, 0, 0, time.Local),
			want:  "2023-08-02",
		},
		{
			input: time.Now(),
			start: time.Date(2023, 07, 20, 0, 0, 0, 0, time.Local),
			end:   time.Date(2023, 07, 24, 0, 0, 0, 0, time.Local),
			want:  "2023-07-27",
		},
		{
			input: time.Date(2023, 07, 28, 0, 0, 0, 0, time.Local),
			start: time.Time{},
			end:   time.Time{},
			want:  "2023-08-02",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test cases-#%d", i), func(t *testing.T) {
			var holidays []string
			for current := tc.start; current.Before(tc.end.AddDate(0, 0, 1)); current = current.AddDate(0, 0, 1) {
				if current.After(tc.input) {
					holidays = append(holidays, current.Format(LayoutDateFormat))
				}
			}

			got := CalculateExpiredAt(tc.input, expirationDays, holidays).Format(LayoutDateFormat)
			if got != tc.want {
				t.Errorf(`want %v, got %v`, tc.want, got)
			}
		})
	}
}

const LayoutDateFormat string = `2006-01-02`

func CalculateExpiredAt(orderDate time.Time, expirationDays int, holidays []string) time.Time {

	expirationDate := orderDate.AddDate(0, 0, expirationDays)
	for current := orderDate.AddDate(0, 0, 1); current.Before(expirationDate.AddDate(0, 0, 1)); current = current.AddDate(0, 0, 1) {
		if IsWeekend(current) || IsHoliday(current, holidays) {
			expirationDate = expirationDate.AddDate(0, 0, 1)
		}
	}

	return expirationDate
}

// IsHoliday validate if time is holiday
func IsHoliday(date time.Time, holidays []string) bool {
	formattedDate := date.Format(LayoutDateFormat)
	for _, holiday := range holidays {
		if holiday == formattedDate {
			return true
		}
	}

	return false
}

// IsWeekend validate if time is weekend
func IsWeekend(date time.Time) bool {
	return date.Weekday() == time.Saturday || date.Weekday() == time.Sunday
}
