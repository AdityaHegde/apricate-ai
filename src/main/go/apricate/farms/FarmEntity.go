package farms

import (
  "github.com/lib/pq"
)

type FarmEntity struct {
  Uuid           string                    `json:"uuid" gorm:"primaryKey"`
  LocationSymbol string                    `json:"location_symbol"`
  Bonuses        pq.StringArray            `json:"bonuses" gorm:"type:text[]"`
  Buildings      map[string]int            `json:"buildings" gorm:"type:jsonb"`
  Plots          map[string]FarmPlotEntity `json:"plots" gorm:"-"`
}
