package day

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"unicode"

	"github.com/yaoapp/kun/log"
)

// Datetime type of day
type Datetime struct {
	time.Time
}

var defaultLocation *time.Location = nil

var defaultFormats = []string{
	time.RFC3339,
	time.RFC3339Nano,
	"2006-01-02T15:04:05-0700",
	"2006-01-02T15:04:05.000Z",
	"2006-01-02T15:04:05+07:00",
	"2006-01-02T15:04:05Z",
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
	"2006-01-02",
	"15:04:05",
}

// Now is a alias of Make
func Now() *Datetime {
	return Make()
}

// Make make a new Datetime
func Make() *Datetime {
	now := &Datetime{Time: time.Now()}
	if defaultLocation != nil {
		now.Time = now.In(defaultLocation)
	}
	return now
}

// Of make a new datetime with the given value
func Of(value interface{}, formats ...string) *Datetime {

	switch value.(type) {
	case time.Time:
		return &Datetime{Time: value.(time.Time)}
	case *Datetime:
		return value.(*Datetime)
	case Datetime:
		d := value.(Datetime)
		return &d
	}

	if len(formats) == 0 {
		formats = defaultFormats
	}
	valueStr := fmt.Sprintf("%v", value)
	for _, format := range formats {
		valueTime, err := time.Parse(format, valueStr)
		if err == nil {
			if defaultLocation != nil {
				valueTime = valueTime.In(defaultLocation)
				return &Datetime{Time: valueTime}
			}
			return &Datetime{Time: valueTime}
		}
		log.Trace("ERROR: %s", err.Error())
	}
	panic("the given value is not time format")
}

// Load load time with given format
func (d *Datetime) Load(value interface{}, formats ...string) *Datetime {
	*d = *Of(value, formats...)
	return d
}

// Timezone set the timezone of datetime
func (d *Datetime) Timezone(name string, offset ...int) *Datetime {
	if len(offset) > 0 {
		loc := time.FixedZone(name, offset[0])
		defaultLocation = loc
		d.Time = d.In(loc)
		return d
	}
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err.Error())
	}
	d.Time = d.In(loc)
	return d
}

// GetTimezone get the timezone of current
func GetTimezone() (name string, offset int) {
	if defaultLocation == nil {
		return time.Now().Zone()
	}
	return time.Now().In(defaultLocation).Zone()
}

// TimezoneSystem using the system default zone
func TimezoneSystem() {
	defaultLocation = nil
}

// TimezoneUTC using the UTC zone
func TimezoneUTC() {
	defaultLocation = time.UTC
}

// Timezone set the default location
func Timezone(name string, offset ...int) {
	if len(offset) > 0 {
		loc := time.FixedZone(name, offset[0])
		defaultLocation = loc
		return
	}

	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err.Error())
	}
	defaultLocation = loc
}

// TimeZones Get a list of valid time zones
func TimeZones() []string {
	var zones []string
	var zoneDirs = []string{
		// Update path according to your OS
		"/usr/share/zoneinfo/",
		"/usr/share/lib/zoneinfo/",
		"/usr/lib/locale/TZ/",
	}

	for _, zd := range zoneDirs {
		zones = walkTzDir(zd, zones)
		for idx, zone := range zones {
			zones[idx] = strings.ReplaceAll(zone, zd+"/", "")
		}
	}

	return zones
}

func walkTzDir(path string, zones []string) []string {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return zones
	}

	isAlpha := func(s string) bool {
		for _, r := range s {
			if !unicode.IsLetter(r) {
				return false
			}
		}
		return true
	}

	for _, info := range fileInfos {
		if info.Name() != strings.ToUpper(info.Name()[:1])+info.Name()[1:] {
			continue
		}

		if !isAlpha(info.Name()[:1]) {
			continue
		}

		newPath := path + "/" + info.Name()

		if info.IsDir() {
			zones = walkTzDir(newPath, zones)
		} else {
			zones = append(zones, newPath)
		}
	}

	return zones
}

// Scan for db scan
func (d *Datetime) Scan(src interface{}) error {
	*d = *Of(src)
	return nil
}

// Value for db driver value
func (d *Datetime) Value() (driver.Value, error) {
	return d.Time, nil
}

// MarshalJSON for json marshalJSON
func (d *Datetime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time)
}

// UnmarshalJSON for json marshalJSON
func (d *Datetime) UnmarshalJSON(data []byte) error {
	*d = *Of(data)
	return nil
}
