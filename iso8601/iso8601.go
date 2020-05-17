package iso8601

import "time"

const Format = "2006-01-02T15:04:05.000Z07"

var fixedZone = time.FixedZone("", 0)

//noinspection GoUnusedExportedFunction
func ZmillisFromString(pTime string) (time.Time, error) {
	return time.ParseInLocation(Format, pTime, fixedZone)
}

func ZmillisToString(pTime *time.Time) string {
	if pTime == nil {
		now := time.Now()
		pTime = &now
	}
	return pTime.UTC().Format(Format)
}

func ZmillisNow() string {
	return ZmillisToString(nil)
}
