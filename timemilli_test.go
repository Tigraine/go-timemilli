package timemilli

import (
	"testing"
	"time"
)

func BenchmarkMult(b *testing.B) {
	var inp int64
	inp = 1542810446506
	for n := 0; n < b.N; n++ {
		time.Unix(0, inp*int64(time.Nanosecond))
	}
}

func BenchmarkDiv(b *testing.B) {
	var ms int64
	ms = 1542810446506
	for n := 0; n < b.N; n++ {
		time.Unix(ms/int64(millisInSecond), (ms%int64(millisInSecond))*int64(time.Nanosecond))
	}
}

func TestFromUnixMilli(t *testing.T) {
	ms := 1542810446506
	ts := FromUnixMilli(int64(ms))
	if ts.UTC().String() != "2018-11-21 14:27:26.506 +0000 UTC" {
		t.Errorf("Expected time to be '2018-11-21 14:27:26.506 +0000 UTC' got '%s'", ts.UTC().String())
	}
}
