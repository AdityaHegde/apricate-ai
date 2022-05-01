package warehouses

type Wares struct {
	Tools   map[string]int `json:"tools" gorm:"type:jsonb"`
	Produce map[string]int `json:"produce" gorm:"type:jsonb"`
	Goods   map[string]int `json:"goods" gorm:"type:jsonb"`
	Seeds   map[string]int `json:"seeds" gorm:"type:jsonb"`
}
