# go-timemilli

The go standard library does not contain a convenient conversion function from Unix milliseconds to `time.Time`.

The only method to convert a unix timestamp is using [`time.Unix`](https://godoc.org/time#Unix) which only takes seconds and nanoseconds. Thus making converting milliseconds a tad too complicated if you don't want to sacrifice accuracy. 

You either have to multiply the millis to nanoseconds or split them into seconds and nanoseconds. This small library (3 LOC) fixes that by providing a convenience function for the conversion.

## Usage

```
import "tigraine/go-timemilli"

msTime := 1542810446506
ts := timemilli.FromUnixMilli(msTime)
```

## Performance Considerations

Shockingly multiplying the millisecond timestamp by 1.000.000 to get nanoseconds is half as slow as turning milliseconds into seconds and then turning the remainder into nanos.

`timemilli_test.go` contains two benchmark tests I used to verify this:

```
goos: darwin
goarch: amd64
pkg: github.com/tigraine/go-timemillis
BenchmarkMult-8         2000000000               0.50 ns/op
BenchmarkDiv-8          2000000000               0.25 ns/op
```
