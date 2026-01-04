package humanizer

import (
	"fmt"
	"math"
)

var sizes = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func HumanizeBytes(s int64, base float64) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}

	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10

	var result string

	if val < 10 {
		if val == float64(int64(val)) {
			result = fmt.Sprintf("%.0f%s", val, suffix)
		} else {
			result = fmt.Sprintf("%.1f%s", val, suffix)
		}
	} else {
		result = fmt.Sprintf("%.0f%s", val, suffix)
	}

	return result
}
