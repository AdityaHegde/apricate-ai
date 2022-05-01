package locations

type PortEntity struct {
  Name              string `json:"name"`
  Symbol            string `json:"symbol" gorm:"primaryKey"`
  ConnectedLocation string `json:"connected_locations"`
  Fare              int    `json:"fare"`
  Duration          int    `json:"duration"`
}
