package patreon

import (
	"github.com/devlikeapro/patrons-perks/internal/patron"
)

type PatreonPlatform struct {
}

func (platform *PatreonPlatform) Load(filePath string) ([]patron.Patron, error) {
	patreonPatrons, err := loadCsvFile(filePath)
	if err != nil {
		return nil, err
	}
	patrons, err := PatreonPatronsToPatrons(patreonPatrons)
	return patrons, err
}
