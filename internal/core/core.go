package core

import "time"

func GetFarInTheFuture() time.Time {
	return time.Date(2199, 1, 1, 0, 0, 0, 0, time.UTC)
}
