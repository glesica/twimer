package main

import (
	"fmt"
	"time"
)

func formatDuration(duration time.Duration) string {
	hours := duration / time.Hour
	minutes := (duration - (hours * time.Hour)) / time.Minute
	seconds := (duration - (hours * time.Hour) - (minutes * time.Minute)) / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
