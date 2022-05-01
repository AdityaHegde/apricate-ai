package markets

type TradeEntity struct {
	LocationSymbol string `json:"location_symbol" gorm:"primaryKey"`
	ItemName       string `json:"item_name" gorm:"primaryKey"`
	ItemCategory   string `json:"item_category"`
	Price          int    `json:"price"`
	Exported       bool   `json:"exported"`
}
