package core

import (
	"fmt"
	. "github.com/devlikeapro/patrons-perks/internal/patron"
	"strings"
	"time"
)

func SaveToDatabase(patrons []Patron, platformName string) {
	for i, patron := range patrons {
		record := savePatron(patron, platformName)
		fmt.Println(i)
		fmt.Println(record)
	}

}

func savePatron(patron Patron, platformName string) PatronRecord {
	now := time.Now()
	return PatronRecord{
		Patron:   patron,
		Platform: strings.ToUpper(platformName),
		Active:   patron.ActiveTill.After(now),
	}
}
