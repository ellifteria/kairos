package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	now := time.Now()
	year := now.Year()
	mon := int(now.Month())
	day := now.Day()
	hour := now.Hour()
	min := now.Minute()
	_, currTimeZoneOffset := now.Zone()
	zone := float64(currTimeZoneOffset) / 60.0 / 60.0

	modePtr := flag.String("format", "short-date-time", "sets format; see https://discord.com/developers/docs/reference#message-formatting-timestamp-styles for more")

	flag.Parse()

	remaining := strings.Join(flag.Args(), " ")

	dateMatcher, err := regexp.Compile(`[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]`)
	check(err)
	timeMatcher, err := regexp.Compile(`[0-9][0-9]:[0-9][0-9]`)
	check(err)
	zoneMatcher, err := regexp.Compile(`UTC(-|\+)([0-9]+\.[0-9]+|[0-9]+)`)
	check(err)

	if dateMatcher.MatchString(remaining) {
		dateStr := dateMatcher.FindString(remaining)
		dateList := strings.Split(dateStr, "-")
		year64, err := strconv.ParseInt(dateList[0], 10, 0)
		check(err)
		year = int(year64)
		mon64, err := strconv.ParseInt(dateList[1], 10, 0)
		check(err)
		mon = int(mon64)
		day64, err := strconv.ParseInt(dateList[2], 10, 0)
		check(err)
		day = int(day64)
	}
	if timeMatcher.MatchString(remaining) {
		timeStr := timeMatcher.FindString(remaining)
		timeList := strings.Split(timeStr, ":")
		hour64, err := strconv.ParseInt(timeList[0], 10, 0)
		check(err)
		hour = int(hour64)
		min64, err := strconv.ParseInt(timeList[1], 10, 0)
		check(err)
		min = int(min64)
	}
	if zoneMatcher.MatchString(remaining) {
		zoneStr := zoneMatcher.FindString(remaining)
		zoneOffsetStr := strings.TrimPrefix(strings.TrimPrefix(zoneStr, "UTC"), "+")
		zoneOffset, err := strconv.ParseFloat(zoneOffsetStr, 64)
		check(err)
		zone = zoneOffset
	}

	var mode string

	switch *modePtr {
	case "short-time":
		mode = "t"
	case "long-time":
		mode = "T"
	case "short-date":
		mode = "d"
	case "long-date":
		mode = "D"
	case "short-date-time":
		mode = "f"
	case "long-datetime":
		mode = "F"
	case "relative":
		mode = "R"
	default:
		panic(fmt.Sprintf("unrecognized timestamp format: %s", *modePtr))
	}

	timeToConvert := time.Date(year, time.Month(mon), day, hour, min, 0, 0, time.FixedZone(fmt.Sprintf("UTC-%02.1f", zone), int(zone*60*60)))

	fmt.Printf("<t:%d:%s>\n", timeToConvert.Unix(), mode)
}
