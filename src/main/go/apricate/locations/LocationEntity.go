package locations

import (
  "github.com/lib/pq"
)

type LocationEntity struct {
  Symbol      string         `json:"symbol" gorm:"primaryKey"`
  Name        string         `json:"name"`
  Description string         `json:"description"`
  IslandName  string         `json:"island_name"`
  X           int            `json:"x"`
  Y           int            `json:"y"`
  NPCs        pq.StringArray `json:"npcs" gorm:"type:text[]"`
}
