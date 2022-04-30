package main

import (
  "AdityaHegde/apricate-ai/src/main/go/apricate/farms"
  "AdityaHegde/apricate-ai/src/main/go/apricate/plants"
  "AdityaHegde/apricate-ai/src/main/go/apricate/users"
  "AdityaHegde/apricate-ai/src/main/go/database"
  "fmt"
)

const HomeLocation = "TS-PR-HF"

func main() {
  db, err := database.Connect()
  if err != nil {
    fmt.Println(err)
    return
  }
  if err != nil {
    fmt.Println(err)
    return
  }

  userService := users.GetUserService(db)
  user, err := userService.GetOrFetch(userService.FieldsService.GetPrimaryValue(nil))
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(user)
  }

  farmService := farms.GetFarmService(db)
  farm, _ := farmService.LoadFromRemote(HomeLocation)
  fmt.Println(len(farm.Plots))

  farmPlotService := farms.GetFarmPlotService(db)
  farmPlots, _ := farmPlotService.LoadAllMapFromRemote("")
  fmt.Println(len(*farmPlots))

  plantService := plants.GetPlantService(db)
  plantEntities, _ := plantService.LoadAllMapFromRemote("")
  fmt.Println(len(*plantEntities))
}
