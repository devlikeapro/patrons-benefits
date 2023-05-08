package patron

import "time"

type Patron struct {
	Platform   string
	Level      string
	Name       string
	Email      string
	Status     string
	ActiveTill time.Time
	Perks      map[string]interface{}
}
