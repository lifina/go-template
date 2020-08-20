package utils

import (
	"time"
)

type Clocker interface {
	Now() time.Time
}

type clocker struct{}

func (c *clocker) Now() time.Time {
	return time.Now()
}

var DefaultClocker Clocker = &clocker{}

var TimeZoneJST = time.FixedZone("Asia/Tokyo", 9*60*60)

func TimeNow() time.Time {
	return DefaultClocker.Now()
}

func TimeJST(t time.Time) time.Time {
	return t.In(TimeZoneJST)
}

func TimeUnix(u int64) time.Time {
	return time.Unix(u, 0).In(TimeZoneJST)
}

func MustParseTime(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}
