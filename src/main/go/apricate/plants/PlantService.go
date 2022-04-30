package plants

import (
  "AdityaHegde/apricate-ai/src/main/go/apricate/commons"
  "gorm.io/gorm"
)

type PlantHandler struct {
  commons.EntityHandler[PlantEntity]
}

func (handler *PlantHandler) EntityLoaded(entity *PlantEntity) {
  for _, growthStage := range entity.GrowthStages {
    growthStage.PlantName = entity.Name
    growthStage.Name = entity.Name + ":" + growthStage.Name
  }
}

func GetPlantService(db *gorm.DB) commons.EntityService[PlantEntity] {
  return commons.EntityService[PlantEntity]{
    Db: db,
    FieldsService: &commons.SimpleEntityFieldsService[PlantEntity]{
      PrimaryKey: "name",
      PrimaryValueGetter: func(entity *PlantEntity) string {
        return entity.Name
      },
    },
    Handler: new(PlantHandler),
    HttpService: commons.HttpService[PlantEntity]{
      UrlService: &commons.SimpleHttpURLService{
        ApiPath: "/plants",
      },
    },
  }
}
