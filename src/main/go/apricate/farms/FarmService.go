package farms

import (
  "AdityaHegde/apricate-ai/src/main/go/apricate/commons"
  "gorm.io/gorm"
)

func GetFarmService(db *gorm.DB) commons.EntityService[FarmEntity] {
  return commons.EntityService[FarmEntity]{
    Db: db,
    FieldsService: &commons.SimpleEntityFieldsService[FarmEntity]{
      PrimaryKey: "uuid",
      PrimaryValueGetter: func(entity *FarmEntity) string {
        return entity.Uuid
      },
    },
    Handler: new(commons.DummyEntityHandler[FarmEntity]),
    HttpService: commons.HttpService[FarmEntity]{
      UrlService: &commons.SimpleHttpURLService{
        ApiPath: "/my/farms",
      },
    },
  }
}
