package actions

import (
	"fmt"
	"math"
)

var sizes = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanizeBytes(s int64, base float64) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}

	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10

	f := "%.0%s"
	if val < 10 {
		f = "%.1f%s"
	}

	return fmt.Sprintf(f, val, suffix)
}

func formatSize(size int64, human bool) string {
	if human {
		return humanizeBytes(size, 1000)
	} else {
		return fmt.Sprintf("%dB", size)
	}
}
