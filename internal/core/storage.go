package core

import (
	"errors"
	. "github.com/devlikeapro/patrons-perks/internal/patron"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
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

type JsonString string

type BenefitRecord struct {
	Patron PatronRecord
	gorm.Model
	// Benefit definition
	Name   string
	Params JsonString
	// Benefit state
	Given     bool
	Skip      bool
	InfoStr   JsonString
	StatusStr JsonString
}

type Storage struct {
	db *gorm.DB
}

func GetStorage() (*Storage, error) {
	db, err := gorm.Open(sqlite.Open("patrons.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&PatronRecord{})
	if err != nil {
		return nil, err
	}
	return &Storage{db}, nil
}

func (storage *Storage) SaveToDatabase(patrons []Patron, platformName string) {
	records := make([]*PatronRecord, 0, len(patrons))
	for _, patron := range patrons {
		record := toPatronRecord(patron, platformName)
		records = append(records, record)
	}
	activePatrons := lo.Filter(records, func(item *PatronRecord, _ int) bool {
		return item.Active
	})

	createdRecords := 0
	for _, record := range records {
		var created bool
		record, created = storage.upsertPatron(record)
		if created {
			createdRecords = createdRecords + 1
		}
	}

	log.WithFields(
		log.Fields{
			"Total":   len(records),
			"Active":  len(activePatrons),
			"Created": createdRecords,
		},
	).Info("Saved records")
}

func (storage *Storage) upsertPatron(patron *PatronRecord) (*PatronRecord, bool) {
	existedPatron := &PatronRecord{}
	result := storage.db.Where("platform = ? and email = ?", patron.Platform, patron.Email).Limit(1).Find(&existedPatron)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result = storage.db.Where("platform = ? and name = ?", patron.Platform, patron.Name).Limit(1).Find(&existedPatron)
	}
	patron.ID = existedPatron.ID
	newRecord := patron.ID == 0
	if newRecord {
		storage.db.Create(&patron)
	} else {
		storage.db.Omit("CreatedAt").Save(&patron)
	}
	return patron, newRecord
}

func toPatronRecord(patron Patron, platformName string) *PatronRecord {
	now := time.Now()
	return &PatronRecord{
		Patron:   patron,
		Platform: strings.ToUpper(platformName),
		Active:   patron.ActiveTill.After(now),
	}
}
