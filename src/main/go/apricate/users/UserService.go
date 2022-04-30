package users

import (
  "AdityaHegde/apricate-ai/src/main/go/apricate/commons"
  "gorm.io/gorm"
  "os"
)

type UserFieldsService struct {
  commons.SimpleEntityFieldsService[UserEntity]
}

func (fieldsService *UserFieldsService) GetPrimaryValue(entity *UserEntity) string {
  return os.Getenv("APRICATE_USER")
}

func GetUserService(db *gorm.DB) commons.EntityService[UserEntity] {
  FieldsService := new(UserFieldsService)
  FieldsService.PrimaryKey = "username"
  return commons.EntityService[UserEntity]{
    Db:            db,
    FieldsService: FieldsService,
    Handler:       new(commons.DummyEntityHandler[UserEntity]),
    HttpService:   GetUserURLService(),
  }
}
