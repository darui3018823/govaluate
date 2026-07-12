// Package datetime contains expression literal date parsing helpers.
package datetime

import "time"

// ISOFormat is the default query serialization format for parsed date literals.
const ISOFormat = "2006-01-02T15:04:05.999999999Z0700"

var formats = [...]string{
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.Kitchen,
	time.RFC3339,
	time.RFC3339Nano,
	"2006-01-02",
	"2006-01-02 15:04",
	"2006-01-02 15:04:05",
	"2006-01-02 15:04:05-07:00",
	"2006-01-02T15Z0700",
	"2006-01-02T15:04Z0700",
	"2006-01-02T15:04:05Z0700",
	ISOFormat,
}

// Parse attempts to parse candidate as one of govaluate's supported date
// literal formats.
func Parse(candidate string) (time.Time, bool) {
	for _, format := range formats {
		if parsed, ok := parseExact(candidate, format); ok {
			return parsed, true
		}
	}

	return time.Time{}, false
}

func parseExact(candidate string, format string) (time.Time, bool) {
	parsed, err := time.ParseInLocation(format, candidate, time.Local)
	if err != nil {
		return time.Time{}, false
	}

	return parsed, true
}
