package patron

import "time"

type Patron struct {
	Level      string
	Name       string
	Email      string
	ActiveTill time.Time
}

type PatronRecord struct {
	Patron   Patron
	Platform string
	Active   bool
}
