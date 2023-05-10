package platforms

import (
	"github.com/devlikeapro/patrons-perks/internal/core"
	"github.com/devlikeapro/patrons-perks/internal/patron"
	boosty_csv "github.com/devlikeapro/patrons-perks/internal/platforms/boostycsv"
)

type Platform interface {
	Load(filePath string) ([]patron.Patron, error)
}

func ImportFromPlatform(platformName string, filePath string) error {
	var platform Platform

	if platformName == "BOOSTY" {
		platform = &boosty_csv.BoostyPlatform{}
	} else {
		panic("Unknown platform" + platformName)
	}

	patrons, err := platform.Load(filePath)
	if err != nil {
		return err
	}
	core.SaveToDatabase(patrons, platformName)
	return nil
}
