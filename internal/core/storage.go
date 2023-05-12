package core

import (
	"errors"
	. "github.com/devlikeapro/patrons-perks/internal/patron"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
	"time"
)

type PatronRecord struct {
	gorm.Model
	Patron
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

	db.AutoMigrate(&PatronRecord{})
	return &Storage{db}, nil
}

func (storage *Storage) SaveToDatabase(patrons []Patron, platformName string) {
	for _, patron := range patrons {
		record := toPatronRecord(patron, platformName)
		record = storage.upsertPatron(record)
	}
}

func (storage *Storage) upsertPatron(patron *PatronRecord) *PatronRecord {
	existed := &PatronRecord{}
	result := storage.db.Where("platform = ? and email = ?", patron.Platform, patron.Email).First(&existed)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result = storage.db.Where("platform = ? and name = ?", patron.Platform, patron.Name).First(&existed)
	}
	patron.ID = existed.ID
	if patron.ID == 0 {
		storage.db.Create(&patron)
	} else {
		storage.db.Omit("CreatedAt").Save(&patron)
	}
	return patron
}

func toPatronRecord(patron Patron, platformName string) *PatronRecord {
	now := time.Now()
	return &PatronRecord{
		Patron:   patron,
		Platform: strings.ToUpper(platformName),
		Active:   patron.ActiveTill.After(now),
	}
}
