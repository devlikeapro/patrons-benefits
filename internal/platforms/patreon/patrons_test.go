package patreon

import (
	"github.com/devlikeapro/patrons-perks/internal/patron"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPatreonPatronsToPatrons(t *testing.T) {
	tests := []struct {
		name            string
		patreonPatrons  []PatreonPatron
		expectedPatrons []patron.Patron
	}{
		// dontouch
		//tech@dontouch.ch
		//Active patron	No	115.87	99	monthly	Pro								2023-04-11 07:40:09.177170	2023-04-11 07:40:11	Paid		91765442	2023-04-11 08:05:15.240916	USD			2023-05-11 00:00:00
		//{"Active patron"
		//},
	}

	for _, test := range tests {
		t.Run(
			test.name,
			func(t *testing.T) {
				patrons, err := PatreonPatronsToPatrons(test.patreonPatrons)
				assert.Nil(t, err)
				assert.Equal(t, test.expectedPatrons, patrons)
			},
		)
	}
}
