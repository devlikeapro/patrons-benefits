package platforms

import (
	"fmt"
	"github.com/devlikeapro/patrons-perks/internal/patron"
	boosty_csv "github.com/devlikeapro/patrons-perks/internal/platforms/boosty-csv"
)

type Platform interface {
	Load(filePath string) ([]patron.Patron, error)
}

func ImportFromPlatform(platformName string, filePath string) error {
	var platform Platform

	if platformName == "BOOSTY" {
		platform = boosty_csv.BoostyPlatform{}
	} else {
		panic("Unknown platform" + platformName)
	}

	patrons, err := platform.Load(filePath)
	if err != nil {
		return err
	}
	fmt.Println(patrons)
	return nil
}
