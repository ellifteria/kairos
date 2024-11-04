package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	currYear := now.Year()
	currMon := int(now.Month())
	currDay := now.Day()
	currHour := now.Hour()
	currMin := now.Minute()
	_, currTimeZoneOffset := now.Zone()
	currTimeZoneOffsetHours := float64(currTimeZoneOffset) / 60.0 / 60.0

	yearPtr := flag.Int("year", currYear, "sets year")
	monPtr := flag.Int("mon", currMon, "sets month")
	dayPtr := flag.Int("day", currDay, "sets day")
	hourPtr := flag.Int("hour", currHour, "sets hour")
	minPtr := flag.Int("min", currMin, "sets min")
	zonePtr := flag.Float64("zone", currTimeZoneOffsetHours, "sets timezone offset from UTC")

	modePtr := flag.String("format", "short-date-time", "sets format; see https://discord.com/developers/docs/reference#message-formatting-timestamp-styles for more")

	flag.Parse()

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

	timeToConvert := time.Date(*yearPtr, time.Month(*monPtr), *dayPtr, *hourPtr, *minPtr, 0, 0, time.FixedZone(fmt.Sprintf("UTC-%02.1f", *zonePtr), int(*zonePtr)*60*60))

	fmt.Printf("<t:%d:%s>", timeToConvert.Unix(), mode)
}
