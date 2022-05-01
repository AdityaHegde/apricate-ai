package assistants

type AssistantEntity struct {
	Uuid             string  `json:"uuid" gorm:"primaryKey"`
	Id               int     `json:"id"`
	Archetype        string  `json:"archetype"`
	Speed            float32 `json:"speed"`
	CarryingCapacity int     `json:"carrying_capacity"`
	Location         string  `json:"location"`
}
