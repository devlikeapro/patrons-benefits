build:
	go build -o patrons

import: build
	./patrons import --file ./data/patreon.csv --platform PATREON
	./patrons import --file ./data/boosty.csv --platform BOOSTY
