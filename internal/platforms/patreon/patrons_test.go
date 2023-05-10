package patreon

import (
	"github.com/devlikeapro/patrons-perks/internal/patron"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPatreonPatronsToPatrons(t *testing.T) {
	tests := []struct {
		name            string
		patreonPatrons  []PatreonPatronRecord
		expectedPatrons []patron.Patron
	}{
		{
			"Active patron",
			[]PatreonPatronRecord{
				{
					Name:           "John",
					Email:          "john@example.com",
					PatronStatus:   "Active patron",
					Tier:           "Plus",
					NextChargeDate: DateTime{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
				},
			},
			[]patron.Patron{
				{
					Level:      "PLUS",
					Name:       "John",
					Email:      "john@example.com",
					ActiveTill: time.Date(2199, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			"Former patron",
			[]PatreonPatronRecord{
				{
					Name:           "John",
					Email:          "john@example.com",
					PatronStatus:   "Former patron",
					Tier:           "Plus",
					NextChargeDate: DateTime{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
				},
			},
			[]patron.Patron{
				{
					Level:      "PLUS",
					Name:       "John",
					Email:      "john@example.com",
					ActiveTill: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			"Declined patron - paid payment",
			[]PatreonPatronRecord{
				{
					Name:             "John",
					Email:            "john@example.com",
					PatronStatus:     "Declined patron",
					LastChargeStatus: "Paid",
					Tier:             "Plus",
					LastChargeDate:   DateTime{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
					NextChargeDate:   DateTime{time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC)},
				},
			},
			[]patron.Patron{
				{
					Level:      "PLUS",
					Name:       "John",
					Email:      "john@example.com",
					ActiveTill: time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			"Declined patron - declined payment",
			[]PatreonPatronRecord{
				{
					Name:             "John",
					Email:            "john@example.com",
					PatronStatus:     "Declined patron",
					LastChargeStatus: "Declined",
					Tier:             "Plus",
					LastChargeDate:   DateTime{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
					NextChargeDate:   DateTime{time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC)},
				},
			},
			[]patron.Patron{
				{
					Level:      "PLUS",
					Name:       "John",
					Email:      "john@example.com",
					ActiveTill: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(
			test.name,
			func(t *testing.T) {
				patrons, err := PatreonPatronsToPatrons(test.patreonPatrons)
				assert.Nil(t, err)
				assert.Equal(t, test.expectedPatrons, patrons)
			},
		)
	}
}
