package database

import (
  "AdityaHegde/apricate-ai/src/main/go/apricate/farms"
  "AdityaHegde/apricate-ai/src/main/go/apricate/locations"
  "AdityaHegde/apricate-ai/src/main/go/apricate/plants"
  "AdityaHegde/apricate-ai/src/main/go/apricate/users"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
  dsn := "host=localhost user=apricate password=apricate dbname=apricate port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, err
  }

  err = db.AutoMigrate(users.UserEntity{})
  if err != nil {
    return nil, err
  }

  err = db.AutoMigrate(plants.GrowthStageEntity{})
  if err != nil {
    return nil, err
  }
  err = db.AutoMigrate(plants.PlantEntity{})
  if err != nil {
    return nil, err
  }

  err = db.AutoMigrate(locations.RegionEntity{})
  if err != nil {
    return nil, err
  }
  err = db.AutoMigrate(locations.IslandEntity{})
  if err != nil {
    return nil, err
  }
  err = db.AutoMigrate(locations.LocationEntity{})
  if err != nil {
    return nil, err
  }

  err = db.AutoMigrate(farms.FarmPlotEntity{})
  if err != nil {
    return nil, err
  }
  err = db.AutoMigrate(farms.FarmEntity{})
  if err != nil {
    return nil, err
  }

  return db, nil
}
