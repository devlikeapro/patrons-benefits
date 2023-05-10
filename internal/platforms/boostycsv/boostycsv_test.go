package boostycsv_test

import (
	"github.com/devlikeapro/patrons-perks/internal/patron"
	"github.com/devlikeapro/patrons-perks/internal/platforms/boostycsv"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestSubscriptionsToPatrons(t *testing.T) {
	tests := []struct {
		name            string
		subscriptions   []boostycsv.BoostySubscriptionRecord
		expectedPatrons []patron.Patron
	}{
		{
			name: "PLUS Subscriber",
			subscriptions: []boostycsv.BoostySubscriptionRecord{
				{
					Name:       "John",
					Email:      "john@example.com",
					Type:       "subscription",
					TotalMoney: rand.Float64(),
					StartDate:  time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
					EndDate:    time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC),
					LevelName:  "Plus",
				},
				{
					Name:       "John",
					Email:      "john@example.com",
					Type:       "subscription",
					TotalMoney: rand.Float64(),
					StartDate:  time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC),
					EndDate:    time.Time{},
					LevelName:  "Plus",
				},
			},
			expectedPatrons: []patron.Patron{
				{
					Level:      "PLUS",
					Name:       "John",
					Email:      "john@example.com",
					ActiveTill: time.Date(2122, 3, 4, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "PLUS Subscriber - unsorted records",
			subscriptions: []boostycsv.BoostySubscriptionRecord{
				{
					Name:       "John",
					Email:      "john@example.com",
					Type:       "subscription",
					TotalMoney: rand.Float64(),
					StartDate:  time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC),
					EndDate:    time.Time{},
					LevelName:  "Plus",
				},
				{
					Name:       "John",
					Email:      "john@example.com",
					Type:       "subscription",
					TotalMoney: rand.Float64(),
					StartDate:  time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
					EndDate:    time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC),
					LevelName:  "Plus",
				},
			},
			expectedPatrons: []patron.Patron{
				{
					Level:      "PLUS",
					Name:       "John",
					Email:      "john@example.com",
					ActiveTill: time.Date(2122, 3, 4, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "PLUS with closed Subscriber",
			subscriptions: []boostycsv.BoostySubscriptionRecord{
				{
					Name:       "John",
					Email:      "john@example.com",
					Type:       "subscription",
					TotalMoney: rand.Float64(),
					StartDate:  time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
					EndDate:    time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC),
					LevelName:  "Plus",
				},
				{
					Name:       "John",
					Email:      "john@example.com",
					Type:       "subscription",
					TotalMoney: rand.Float64(),
					StartDate:  time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC),
					EndDate:    time.Date(2022, 4, 4, 0, 0, 0, 0, time.UTC),
					LevelName:  "Plus",
				},
			},
			expectedPatrons: []patron.Patron{
				{
					Level:      "PLUS",
					Name:       "John",
					Email:      "john@example.com",
					ActiveTill: time.Date(2022, 4, 4, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "Used be PLUS, not just a follower",
			subscriptions: []boostycsv.BoostySubscriptionRecord{
				{
					Name:       "John",
					Email:      "john@example.com",
					Type:       "subscription",
					TotalMoney: rand.Float64(),
					StartDate:  time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
					EndDate:    time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC),
					LevelName:  "Plus",
				},
				{
					Name:       "John",
					Email:      "john@example.com",
					Type:       "following",
					TotalMoney: rand.Float64(),
					StartDate:  time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC),
					EndDate:    time.Time{},
					LevelName:  "Follower",
				},
			},
			expectedPatrons: []patron.Patron{
				{
					Level:      "PLUS",
					Name:       "John",
					Email:      "john@example.com",
					ActiveTill: time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "Just a follower",
			subscriptions: []boostycsv.BoostySubscriptionRecord{
				{
					Name:       "John",
					Email:      "john@example.com",
					Type:       "following",
					TotalMoney: rand.Float64(),
					StartDate:  time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC),
					EndDate:    time.Time{},
					LevelName:  "Follower",
				},
			},
			expectedPatrons: []patron.Patron{},
		},
	}

	for _, test := range tests {
		t.Run(
			test.name,
			func(t *testing.T) {
				patrons, err := boostycsv.SubscriptionsToPatrons(test.subscriptions)
				assert.Nil(t, err)
				assert.Equal(t, test.expectedPatrons, patrons)
			},
		)
	}
}
