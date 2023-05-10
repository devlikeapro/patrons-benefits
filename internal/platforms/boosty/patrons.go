package boosty

import (
	"github.com/devlikeapro/patrons-perks/internal/core"
	"github.com/devlikeapro/patrons-perks/internal/patron"
	"github.com/samber/lo"
	"sort"
	"strings"
	"time"
)

// BoostySubscriptionRecord pure records from CSV file
type BoostySubscriptionRecord struct {
	Name       string
	Email      string
	Type       string
	Price      int
	TotalMoney float64
	StartDate  time.Time
	EndDate    time.Time
	LevelName  string
}

const (
	Follower string = "Follower"
)

func SubscriptionsToPatrons(subscriptions []BoostySubscriptionRecord) ([]patron.Patron, error) {
	subscriptions = onlyWithTiers(subscriptions[:])
	recordsByEmail := lo.GroupBy(subscriptions, func(item BoostySubscriptionRecord) string {
		return item.Email
	})

	patrons := make([]patron.Patron, 0, 0)
	for _, records := range recordsByEmail {
		thePatron := getPatron(records)
		patrons = append(patrons, thePatron)
	}
	return patrons, nil
}

func onlyWithTiers(subscriptions []BoostySubscriptionRecord) []BoostySubscriptionRecord {
	return lo.Filter(subscriptions, func(item BoostySubscriptionRecord, i int) bool {
		return item.LevelName != Follower
	})
}

func sortByEndDate(subscriptions []BoostySubscriptionRecord) {
	// subscriptions must be sorted by EndDate - actual last
	sort.Slice(subscriptions, func(i int, j int) bool {
		if subscriptions[i].EndDate.IsZero() {
			return false
		}
		if subscriptions[j].EndDate.IsZero() {
			return true
		}
		return subscriptions[i].EndDate.Before(subscriptions[j].EndDate)
	})
}

// getPatron returns a single Patron for subscripts for SINGLE patron
func getPatron(subscriptions []BoostySubscriptionRecord) patron.Patron {
	sortByEndDate(subscriptions)
	length := len(subscriptions)
	last := subscriptions[length-1]
	var activeTill time.Time
	if last.EndDate.IsZero() {
		activeTill = core.GetFarInTheFuture()
	} else {
		activeTill = last.EndDate
	}

	return patron.Patron{
		Level:      strings.ToUpper(last.LevelName),
		Name:       last.Name,
		Email:      last.Email,
		ActiveTill: activeTill,
	}
}
