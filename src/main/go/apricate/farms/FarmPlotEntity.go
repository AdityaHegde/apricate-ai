package farms

var PlotSize = map[string]int{
  "Miniature": 1,
  "Tiny":      2,
  "Small":     4,
  "Modest":    8,
  "Average":   16,
  "Large":     32,
  "Huge":      64,
}

type PlantedCrop struct {
  Name          string  `json:"name"`
  Size          string  `json:"size"`
  CurrentStage  int     `json:"current_stage"`
  YieldModifier float32 `json:"yield_modifier"`
}

type FarmPlotEntity struct {
  Uuid                    string      `json:"uuid" gorm:"primaryKey"`
  LocationSymbol          string      `json:"location_symbol"`
  Id                      int         `json:"id"`
  Size                    string      `json:"size"`
  PlantQuantity           int         `json:"plant_quantity"`
  GrowthCompleteTimestamp int         `json:"growth_complete_timestamp"`
  Plant                   PlantedCrop `json:"plant" gorm:"type:jsonb"`
}
