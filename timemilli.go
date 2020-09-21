package timemilli

import "time"

const millisInSecond = 1000
const microsInSecond = 1000000
const nsInSecond = 1000000

// Converts Unix Epoch from milliseconds to time.Time
func FromUnixMilli(ms int64) time.Time {
	return time.Unix(ms/int64(millisInSecond), (ms%int64(millisInSecond))*int64(nsInSecond))
}

func FromUnixMicro(ms int64) time.Time {
	return time.Unix(ms/int64(microsInSecond), (ms%int64(microsInSecond))*int64(nsInSecond))
}
