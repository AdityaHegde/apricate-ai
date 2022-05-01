package plants

import (
	"AdityaHegde/apricate-ai/src/main/go/apricate/commons"
	"gorm.io/gorm"
)

type PlantService struct {
  commons.EntityService[PlantEntity]
}

type PlantHandler struct {
  commons.EntityHandler[PlantEntity]
}

func (handler *PlantHandler) EntityLoaded(entity *PlantEntity) {
  for _, growthStage := range entity.GrowthStages {
    growthStage.PlantName = entity.Name
    growthStage.Name = entity.Name + ":" + growthStage.Name
  }
}

func NewPlantService(db *gorm.DB) PlantService {
  plantService := PlantService{}
  plantService.Db = db
  plantService.FieldsService = &commons.SimpleEntityFieldsService[PlantEntity]{
    PrimaryKey: "name",
    PrimaryValueGetter: func(entity *PlantEntity) string {
      return entity.Name
    },
  }
  plantService.Handler = &commons.SimpleEntityHandler[PlantEntity]{
    LoadedHandler: func(entity *PlantEntity) {
      for index, growthStage := range entity.GrowthStages {
        growthStage.PlantName = entity.Name
        growthStage.Index = index
        if growthStage.ConsumableOptions != nil {
          for _, consumableOption := range growthStage.ConsumableOptions {
            consumableOption.PlantName = entity.Name
            consumableOption.GrowthStage = growthStage.Name
          }
        }
      }
    },
  }
  plantService.HttpService = commons.HttpService[PlantEntity]{
    UrlService: &commons.SimpleHttpURLService{
      ApiPath: "/plants",
    },
  }
  return plantService
}
