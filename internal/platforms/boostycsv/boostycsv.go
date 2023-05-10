package boostycsv

import (
	"encoding/csv"
	"fmt"
	"github.com/devlikeapro/patrons-perks/internal/patron"
	"github.com/samber/lo"
	"os"
	"sort"
	"strings"
	"time"
)

type BoostySubscription struct {
	Name       string
	Email      string
	Type       string
	Price      int
	TotalMoney float64
	StartDate  time.Time
	EndDate    time.Time
	LevelName  string
}
type BoostyPlatform struct {
}

const (
	Follower string = "Follower"
)

func (platform *BoostyPlatform) Load(filePath string) ([]patron.Patron, error) {
	subscriptions, err := loadCsvFile(filePath)
	if err != nil {
		return nil, err
	}
	patrons, err := SubscriptionsToPatrons(subscriptions)
	return patrons, err
}

func SubscriptionsToPatrons(subscriptions []BoostySubscription) ([]patron.Patron, error) {
	subscriptions = onlyWithTiers(subscriptions[:])
	recordsByEmail := lo.GroupBy(subscriptions, func(item BoostySubscription) string {
		return item.Email
	})

	patrons := make([]patron.Patron, 0, 0)
	for _, records := range recordsByEmail {
		thePatron := getPatron(records)
		patrons = append(patrons, thePatron)
	}
	return patrons, nil
}

func onlyWithTiers(subscriptions []BoostySubscription) []BoostySubscription {
	return lo.Filter(subscriptions, func(item BoostySubscription, i int) bool {
		return item.LevelName != Follower
	})
}

func sortByEndDate(subscriptions []BoostySubscription) {
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
func getPatron(subscriptions []BoostySubscription) patron.Patron {
	sortByEndDate(subscriptions)
	length := len(subscriptions)
	last := subscriptions[length-1]
	var activeTill time.Time
	if last.EndDate.IsZero() {
		activeTill = last.StartDate.AddDate(0, 1, 1)
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

func loadCsvFile(filePath string) ([]BoostySubscription, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	subscriptions := make([]BoostySubscription, 0, len(records)-1)

	for i, record := range records {
		if i == 0 {
			continue // skip header row
		}

		startDate, err := time.Parse(time.DateOnly, record[5])
		if err != nil {
			return nil, err
		}

		var endDate time.Time
		if record[6] != "-" {
			endDate, err = time.Parse("2006-01-02", record[6])
			if err != nil {
				return nil, err
			}
		}

		price := 0
		fmt.Sscanf(record[3], "%d", &price)

		totalMoney := 0.0
		fmt.Sscanf(record[4], "%f", &totalMoney)

		subscription := BoostySubscription{
			Name:       record[0],
			Email:      record[1],
			Type:       record[2],
			Price:      price,
			TotalMoney: totalMoney,
			StartDate:  startDate,
			EndDate:    endDate,
			LevelName:  record[7],
		}

		subscriptions = append(subscriptions, subscription)
	}
	return subscriptions, nil
}
