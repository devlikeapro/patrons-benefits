package patron

import "time"

type Patron struct {
	Platform   string
	Level      string
	Name       string
	Email      string
	Active     bool
	ActiveTill time.Time
	Perks      map[string]interface{}
}
