// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
	"time"
)

const (
	// LayoutDateTimeFormat date time format
	LayoutDateTimeFormat = `2006-01-02 03:04:05`

	// TimeZoneAsiaJakarta format
	TimeZoneAsiaJakarta = `Asia/Jakarta`
)

func main() {
	// var cstJkt, _ = time.LoadLocation("Asia/Jakarta")
	// fmt.Println("JKT : ", time.Now().In(cstJkt).Format("2006-01-02 15:04:05"))
	result := "2022-06-29 21:40:00"

	tn := time.Now().Format(LayoutDateTimeFormat)
	isValid, _ := IsValidExecutionTime(time.Minute*10, result, tn)
	fmt.Println(isValid)
	fmt.Println(tn)

	expiredAt := time.Now().Add(10 * time.Hour)
	ttlCache := TimeDuration(expiredAt, time.Now())

	if ttlCache.Minutes() > 60 {
		fmt.Println(RoundFloat(ttlCache.Hours(), 0), "hour")
	}

	if ttlCache.Seconds() > 60 {
		fmt.Println(RoundFloat(ttlCache.Minutes(), 0), "minute")
	}

	fmt.Println(RoundFloat(ttlCache.Seconds(), 0), "second")
}

// IsValidExecutionTime validate execution time
func IsValidExecutionTime(eet time.Duration, createdTime, nowTime string) (bool, error) {
	it, err := time.Parse(LayoutDateTimeFormat, createdTime)

	if err != nil {
		return false, fmt.Errorf("initiated time error %v ", err)
	}

	expire := it.Add(eet)

	nt, err := time.Parse(LayoutDateTimeFormat, nowTime)
	if err != nil {
		return false, fmt.Errorf("time now error %v ", err)
	}

	if nt.Unix() >= expire.Unix() {
		return false, nil
	}

	return true, nil
}

func TimeDuration(timeAfter, timeBefore time.Time) time.Duration {
	return timeAfter.Sub(timeBefore)
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
