package users

import (
  "github.com/lib/pq"
  "gorm.io/datatypes"
)

type UserLedger struct {
  Currencies datatypes.JSONMap `json:"currencies" gorm:"type:jsonb"`
  Favor      datatypes.JSONMap `json:"favor" gorm:"type:jsonb"`
}

type UserEntity struct {
  Username     string         `json:"username" gorm:"primaryKey"`
  Title        string         `json:"title"`
  UserSince    int            `json:"user-since"`
  Ledger       UserLedger     `json:"ledger" gorm:"embedded"`
  Achievements pq.StringArray `json:"achievements" gorm:"type:text[]"`
  Contracts    pq.StringArray `json:"contracts" gorm:"type:text[]"`
  Assistants   pq.StringArray `json:"assistants" gorm:"type:text[]"`
  Farms        pq.StringArray `json:"farms" gorm:"type:text[]"`
  Plots        pq.StringArray `json:"plots" gorm:"type:text[]"`
  Warehouses   pq.StringArray `json:"warehouses" gorm:"type:text[]"`
}
