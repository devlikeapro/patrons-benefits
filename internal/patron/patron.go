package patron

import "time"

type Patron struct {
	Level      string
	Name       string
	Email      string
	ActiveTill time.Time
}
