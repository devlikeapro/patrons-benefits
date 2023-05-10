package boosty

import "github.com/devlikeapro/patrons-perks/internal/patron"

type BoostyPlatform struct {
}

func (platform *BoostyPlatform) Load(filePath string) ([]patron.Patron, error) {
	subscriptions, err := loadCsvFile(filePath)
	if err != nil {
		return nil, err
	}
	patrons, err := SubscriptionsToPatrons(subscriptions)
	return patrons, err
}
