package users

type UserEntity struct {
  Username     string   `json:"username" gorm:"primaryKey"`
  Title        string   `json:"title"`
  UserSince    int      `json:"user-since"`
  Achievements []string `json:"achievements"`
  Contracts    []string `json:"contracts"`
  Assistants   []string `json:"assistants"`
  Farms        []string `json:"farms"`
  Plots        []string `json:"plots"`
  Warehouses   []string `json:"warehouses"`
}
