package commons

import "strings"

type SimpleHttpURLService struct {
  HttpURLService
  ApiPath string
}

func (service *SimpleHttpURLService) GetSingleUrl(key string) string {
  return service.ApiPath + "/" + strings.ReplaceAll(key, " ", "_")
}
func (service *SimpleHttpURLService) GetManyUrl(key string) string {
  return service.ApiPath
}
