package patreon

import (
	"fmt"
	"github.com/devlikeapro/patrons-perks/internal/core"
	"github.com/devlikeapro/patrons-perks/internal/patron"
	"strings"
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

const (
	DeclinedPatron = "Declined patron"
	ActivePatron   = "Active patron"
	FormerPatron   = "Former patron"
)

const (
	LastChargeStatusDeclined = "Declined"
	LastChargeStatusPaid     = "Paid"
)

type PatreonPatronRow struct {
	Name             string   `csv:"Name"`
	Email            string   `csv:"Email"`
	LastChargeDate   DateTime `csv:"Last Charge Date"`
	LastChargeStatus string   `csv:"Last Charge Status"`
	PatronStatus     string   `csv:"Patron Status"`
	Tier             string   `csv:"Tier"`
	NextChargeDate   DateTime `csv:"Next Charge Date"`
}

func PatreonPatronsToPatrons(patreonPatrons []PatreonPatronRow) ([]patron.Patron, error) {
	patrons := make([]patron.Patron, 0, len(patreonPatrons))
	for _, patreonPatron := range patreonPatrons {
		var activeTill time.Time
		switch patreonPatron.PatronStatus {
		case ActivePatron:
			activeTill = core.GetFarInTheFuture()
		case FormerPatron:
			activeTill = patreonPatron.NextChargeDate.Time
		case DeclinedPatron:
			if patreonPatron.LastChargeStatus == LastChargeStatusPaid {
				// They have paid, so they can use the rest
				activeTill = patreonPatron.NextChargeDate.Time
			} else {
				// They haven't paid, so they can't use benefits
				activeTill = patreonPatron.LastChargeDate.Time
			}
		case "":
			continue
		default:
			err := fmt.Errorf("unknown patreon status - %s", patreonPatron.PatronStatus)
			return nil, err
		}

		apatron := patron.Patron{
			Level:      strings.ToUpper(patreonPatron.Tier),
			Name:       patreonPatron.Name,
			Email:      patreonPatron.Email,
			ActiveTill: activeTill,
		}
		patrons = append(patrons, apatron)

	}
	return patrons, nil
}
