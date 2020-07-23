package timemilli

import "time"

const millisInSecond = 1000
const nsInMilliSecond = 1000_000

// Converts Unix Epoch from milliseconds to time.Time
func FromUnixMilli(ms int64) time.Time {
	return time.Unix(ms/int64(millisInSecond), (ms%int64(millisInSecond))*int64(nsInMilliSecond))
}
