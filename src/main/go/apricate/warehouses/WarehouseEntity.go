package warehouses

type WarehouseEntity struct {
  Uuid           string         `json:"uuid" gorm:"primaryKey"`
  LocationSymbol string         `json:"location_symbol"`
  Tools          map[string]int `json:"tools" gson:"type:jsonb"`
  Produce        map[string]int `json:"produce" gorm:"type:jsonb"`
  Goods          map[string]int `json:"goods" gorm:"type:jsonb"`
  Seeds          map[string]int `json:"seeds" gorm:"type:jsonb"`
}
