package markets

import (
	"AdityaHegde/apricate-ai/src/main/go/apricate/warehouses"
)

type MarketEntity struct {
	Name           string           `json:"name"`
	LocationSymbol string           `json:"location_symbol" gorm:"primaryKey"`
	Imports        warehouses.Wares `json:"imports" gorm:"type:jsonb"`
	Exports        warehouses.Wares `json:"exports" gorm:"type:jsonb"`
}
