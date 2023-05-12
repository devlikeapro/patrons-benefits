package patreon

import (
	"github.com/gocarina/gocsv"
	"os"
)

func loadCsvFile(filePath string) ([]PatreonPatronRow, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var patrons []PatreonPatronRow
	err = gocsv.UnmarshalFile(file, &patrons)
	if err != nil {
		return nil, err
	}
	return patrons, nil
}
