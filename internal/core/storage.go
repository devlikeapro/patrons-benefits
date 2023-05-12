package core

import (
	. "github.com/devlikeapro/patrons-perks/internal/patron"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
	"time"
)

type PatronRecord struct {
	gorm.Model
	Patron   Patron
	Platform string
	Active   bool
}

type Storage struct {
	db *gorm.DB
}

func GetStorage() (*Storage, error) {
	db, err := gorm.Open(sqlite.Open("patrons.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Storage{db}, nil
}

func (storage *Storage) SaveToDatabase(patrons []Patron, platformName string) {
	rows := make([]PatronRecord, 0, len(patrons))
	for _, patron := range patrons {
		record := toPatronRecord(patron, platformName)
		rows = append(rows, record)
	}
	storage.db.Create(rows)
}

func toPatronRecord(patron Patron, platformName string) PatronRecord {
	now := time.Now()
	return PatronRecord{
		Patron:   patron,
		Platform: strings.ToUpper(platformName),
		Active:   patron.ActiveTill.After(now),
	}
}
