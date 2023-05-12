package benefits

import (
	"github.com/devlikeapro/patrons-perks/internal/core"
)

// BenefitPatronInfo contains additional information about the patron's benefit
// like username in third party system or id
type BenefitPatronInfo interface {
}

// BenefitResult contains the result status - generator password id or other information
// about ADDED benefit
type BenefitResult interface {
}

type Benefit interface {
	GiveTo(patron *core.PatronRecord, info *BenefitPatronInfo, result *BenefitResult) (*BenefitResult, error)
	TakeFrom(patron *core.PatronRecord, info *BenefitPatronInfo, result *BenefitResult) (*BenefitResult, error)
	GetBenefitName() string
	ParseBenefitInfo(string string) *BenefitPatronInfo
	ParseBenefitResult(string string) *BenefitResult
}

type BenefitsForLevels map[string][]*Benefit

func (benefitsForLevels BenefitsForLevels) ProcessActive(patron *core.PatronRecord) {
	benefits := benefitsForLevels[patron.Level]
	for _, benefit := range benefits {
		benefit.GiveTo(patron)
	}
}
