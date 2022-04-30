package farms

import (
  "AdityaHegde/apricate-ai/src/main/go/apricate/commons"
  "gorm.io/gorm"
)

func GetFarmPlotService(db *gorm.DB) commons.EntityService[FarmPlotEntity] {
  return commons.EntityService[FarmPlotEntity]{
    Db: db,
    FieldsService: &commons.SimpleEntityFieldsService[FarmPlotEntity]{
      PrimaryKey: "uuid",
      PrimaryValueGetter: func(entity *FarmPlotEntity) string {
        return entity.Uuid
      },
    },
    Handler: new(commons.DummyEntityHandler[FarmPlotEntity]),
    HttpService: commons.HttpService[FarmPlotEntity]{
      UrlService: &commons.SimpleHttpURLService{
        ApiPath: "/my/plots",
      },
    },
  }
}
