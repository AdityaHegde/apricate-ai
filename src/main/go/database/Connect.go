package database

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
  dsn := "host=localhost user=apricate password=apricate dbname=apricate port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, err
  }

  return db, nil
}
