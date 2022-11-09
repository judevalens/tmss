package misc

const utcDiff = (70*365 + 17) * 24 * 60 * 60

func NtpToUtc(ts uint64) int64 {
	seconds := (ts >> 32) & 0xFFFFFFFF
	milliSeconds := ts & 0xFFFFFFFF
	u := seconds - utcDiff
	u += milliSeconds / 1000
	return int64(u)
}

func UtcToNtp(ts int64) uint64 {
	var seconds uint64
	seconds = seconds & uint64(ts)
	seconds <<= 32
	return seconds
}
