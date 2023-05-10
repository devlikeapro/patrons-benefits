package boosty

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func loadCsvFile(filePath string) ([]BoostySubscriptionRecord, error) {
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

	subscriptions := make([]BoostySubscriptionRecord, 0, len(records)-1)

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

		subscription := BoostySubscriptionRecord{
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
