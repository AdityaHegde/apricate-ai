package plants

type ConsumableOption struct {
	PlantName   string  `gorm:"primaryKey;foreignKey"`
	GrowthStage string  `gorm:"primaryKey;foreignKey"`
	Name        string  `json:"name" gorm:"primaryKey"`
	Quantity    int     `json:"quantity"`
	AddedYield  float32 `json:"addedYield"`
}

type Harvestable struct {
	Produce      map[string]float32 `json:"produce" gorm:"type:jsonb"`
	Seeds        map[string]float32 `json:"seeds" gorm:"type:jsonb"`
	Goods        map[string]float32 `json:"goods" gorm:"type:jsonb"`
	FinalHarvest bool               `json:"final_harvest"`
}

type GrowthStageEntity struct {
	Index             int                `gorm:"index"`
	Name              string             `json:"name" gorm:"primaryKey"`
	PlantName         string             `gorm:"primaryKey;foreignKey"`
	Description       string             `json:"description"`
	Action            string             `json:"action"`
	Skippable         bool               `json:"skippable"`
	Repeatable        bool               `json:"repeatable"`
	ConsumableOptions []ConsumableOption `json:"consumable_options" gorm:"foreignKey:PlantName,GrowthStage"`
	AddedYield        float32            `json:"added_yield"`
	GrowthTime        int                `json:"growth_time"`
	Harvestable       Harvestable        `json:"harvestable" gorm:"type:jsonb"`
}

type PlantEntity struct {
	Name         string              `json:"name" gorm:"primaryKey"`
	Description  string              `json:"description"`
	MinSize      string              `json:"min_size"`
	MaxSize      string              `json:"max_size"`
	GrowthStages []GrowthStageEntity `json:"growth_stages" gorm:"foreignKey:PlantName"`
}
