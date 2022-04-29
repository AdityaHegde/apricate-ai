package database

import (
  "AdityaHegde/apricate-ai/src/main/go/apricate-clients/users"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
  dsn := "host=localhost user=apricate password=apricate dbname=apricate port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, err
  }

  db.AutoMigrate(users.UserEntity{})

  return db, nil
}
