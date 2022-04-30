package users

import "AdityaHegde/apricate-ai/src/main/go/apricate/commons"

type UserURLService struct {
  commons.HttpURLService
}

func (urlService *UserURLService) GetSingleUrl(key string) string {
  return "/my/user"
}
func (urlService *UserURLService) GetManyUrl(key string) string {
  return ""
}

func GetUserURLService() commons.HttpService[UserEntity] {
  return commons.HttpService[UserEntity]{
    UrlService: new(UserURLService),
  }
}
