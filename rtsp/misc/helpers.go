package misc

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func BuildIsoTime(startTimeInput string) *time.Time {
	isoTimeExpr := regexp.MustCompile("[:.]")
	timeParts := isoTimeExpr.Split(startTimeInput, -1)
	var startTime time.Time

	if len(timeParts) == 3 {
		startTime = time.Date(0, 0, 0, parseInt(timeParts[0]), parseInt(timeParts[1]), parseInt(timeParts[2]), 0, time.Local)
	} else if len(timeParts) == 4 {
		float, err := strconv.ParseFloat("."+timeParts[3], 32)
		if err != nil {
			return nil
		}
		tNano := float * float64(time.Second.Nanoseconds())
		startTime = time.Date(0, 0, 0, parseInt(timeParts[0]), parseInt(timeParts[1]), parseInt(timeParts[2]), int(math.Ceil(tNano)), time.Local)
	}
	return &startTime
}

func parseInt(input string) int {
	n, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
		return 0
	}
	return n

}

func TimeToSec(t *time.Time) float64 {
	tSec := (float64(t.Hour()) * 60 * 60) + (float64(t.Minute()) * 60) + float64(t.Second()) + float64(t.Nanosecond())/(1000*1000*1000)
	return math.Round(tSec*1000) / 1000
}

// Removes the start of line and end of line anchors in a sub regex
func RegexToString(regex *regexp.Regexp) string {
	r := strings.TrimPrefix(regex.String(), "^")
	return strings.TrimSuffix(r, "$")
}
