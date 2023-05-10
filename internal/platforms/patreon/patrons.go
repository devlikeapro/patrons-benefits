package patreon

import (
	"github.com/devlikeapro/patrons-perks/internal/patron"
	"time"
)

type DateTime struct {
	time.Time
}

func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	if csv == "" {
		date.Time = time.Time{}
		return nil
	}
	date.Time, err = time.Parse(time.DateTime, csv)
	return err
}

type PatreonPatron struct {
	Name              string   `csv:"Name"`
	MaxPosts          int      `csv:"Max Posts"`
	Email             string   `csv:"Email"`
	PatronStatus      string   `csv:"Patron Status"`
	Tier              string   `csv:"Tier"`
	LastChargeStatus  string   `csv:"Last Charge Status"`
	AdditionalDetails string   `csv:"Additional Details"`
	LastUpdated       DateTime `csv:"Last Updated"`
	AccessExpiration  string   `csv:"Access Expiration"`
	NextChargeDate    DateTime `csv:"Next Charge Date"`
}

func PatreonPatronsToPatrons(patrons []PatreonPatron) ([]patron.Patron, error) {
	return []patron.Patron{}, nil
}
