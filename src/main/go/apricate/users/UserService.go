package users

import (
	"AdityaHegde/apricate-ai/src/main/go/apricate/commons"
	"gorm.io/gorm"
	"os"
)

type UserService struct {
	commons.EntityService[UserEntity]
}

type UserFieldsService struct {
	commons.SimpleEntityFieldsService[UserEntity]
}

func (fieldsService *UserFieldsService) GetPrimaryValue(entity *UserEntity) string {
	return os.Getenv("APRICATE_USER")
}

func NewUserService(db *gorm.DB) UserService {
	userService := UserService{}
	userService.Db = db
	FieldsService := new(UserFieldsService)
	FieldsService.PrimaryKey = "username"
	userService.FieldsService = FieldsService
	userService.Handler = new(commons.DummyEntityHandler[UserEntity])
	userService.HttpService = NewUserURLService()
	return userService
}
