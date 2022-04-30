package locations

type RegionIsland struct {
  Symbol string `json:"symbol"`
  Name   string `json:"name"`
}

type RegionEntity struct {
  Symbol      string         `json:"symbol" gorm:"primaryKey"`
  Name        string         `json:"name"`
  RegionGroup string         `json:"region_group"`
  Description string         `json:"description"`
  Islands     []RegionIsland `json:"islands" gorm:"type:jsonb"`
}
