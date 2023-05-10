package boosty_csv

import (
	"encoding/csv"
	"fmt"
	"github.com/devlikeapro/patrons-perks/internal/patron"
	"os"
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

func (platform *BoostyPlatform) Load(filePath string) ([]patron.Patron, error) {
	subscriptions, err := loadCsvFile(filePath)
	if err != nil {
		return nil, err
	}
	fmt.Println(subscriptions)
	patrons, err := subscriptionsToPatrons(subscriptions)
	return patrons, err
}

func subscriptionsToPatrons(subscriptions []BoostySubscription) ([]patron.Patron, error) {
	patrons := make([]patron.Patron, 0, 0)
	return patrons, nil

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

		startDate, err := time.Parse("2006-01-02", record[5])
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
