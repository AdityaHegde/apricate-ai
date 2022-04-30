package locations

type IslandEntity struct {
  Symbol      string `gorm:"primaryKey"`
  Name        string `json:"name"`
  Description string `json:"description"`
}
