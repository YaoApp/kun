package day

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Day type of day
type Day struct {
	value  interface{}
	locale string
}

// Of make a new day
func Of(value interface{}) Day {
	return Day{value: value, locale: "default"}
}

// Now make a new day of now
func Now() Day {
	return Day{value: time.Now(), locale: "default"}
}

// Time cast the Day to time.Time
func (d Day) Time(formats ...string) time.Time {
	return time.Now()
}

// Locale set the locale of day
func (d Day) Locale(name string) Day {
	d.locale = name
	return d
}

// Format cast to string with format
func (d Day) Format(format string) string {
	return "2021-05-18 12:30:20"
}

// Human cast to human language
func (d Day) Human() string {
	return "1 hour ago"
}

// Scan for db scan
func (d *Day) Scan(src interface{}) error {
	*d = Of(src)
	return nil
}

// Value for db driver value
func (d *Day) Value() (driver.Value, error) {
	return d.Time(), nil
}

// MarshalJSON for json marshalJSON
func (d *Day) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.value)
}

// UnmarshalJSON for json marshalJSON
func (d *Day) UnmarshalJSON(data []byte) error {
	*d = Of(data)
	return nil
}
