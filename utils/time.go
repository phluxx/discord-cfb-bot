package utils

import (
	"time"
)

func ParseGameTime(gameTime string) (time.Time, error) {
	return time.Parse("03:04 PM", gameTime)
}
