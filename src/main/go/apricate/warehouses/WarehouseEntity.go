package warehouses

type WarehouseEntity struct {
	Wares
	Uuid           string `json:"uuid" gorm:"primaryKey"`
	LocationSymbol string `json:"location_symbol"`
}
