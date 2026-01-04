package actions

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

func formatSize(size int64, human bool) string {
	if human {
		return humanize.Bytes(uint64(size))
	} else {
		return fmt.Sprintf("%d B", size)
	}
}
