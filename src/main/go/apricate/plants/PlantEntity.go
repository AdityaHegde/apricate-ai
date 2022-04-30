package plants

type ConsumableOption struct {
  Name       string  `json:"name"`
  Quantity   int     `json:"quantity"`
  AddedYield float32 `json:"addedYield"`
}

type Harvestable struct {
  Produce      map[string]float32 `json:"produce" gson:"type:jsonb"`
  Seeds        map[string]float32 `json:"seeds" gson:"type:jsonb"`
  Goods        map[string]float32 `json:"goods" gson:"type:jsonb"`
  FinalHarvest bool               `json:"final_harvest"`
}

type GrowthStageEntity struct {
  Name              string             `json:"name" gorm:"primaryKey"`
  PlantName         string             `gorm:"foreignKey"`
  Description       string             `json:"description"`
  Action            string             `json:"action"`
  Skippable         bool               `json:"skippable"`
  Repeatable        bool               `json:"repeatable"`
  ConsumableOptions []ConsumableOption `json:"consumable_options" gorm:"-"`
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
