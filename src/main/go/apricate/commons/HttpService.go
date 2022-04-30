package commons

type HttpService[Entity interface{}] struct {
  UrlService HttpURLService
}
type HttpURLService interface {
  GetSingleUrl(key string) string
  GetManyUrl(key string) string
}

func (service *HttpService[Entity]) GetRemoteSingle(key string) (*Entity, error) {
  var entityResp ApricateResponse[Entity]
  _, err := HttpClientGet(&entityResp, service.UrlService.GetSingleUrl(key))
  if err != nil {
    return nil, err
  }
  return &entityResp.Data, nil
}

func (service *HttpService[Entity]) GetRemoteMany(key string) (*[]Entity, error) {
  var entityResp ApricateResponse[[]Entity]
  _, err := HttpClientGet(&entityResp, service.UrlService.GetManyUrl(key))
  if err != nil {
    return nil, err
  }
  return &entityResp.Data, nil
}

func (service *HttpService[Entity]) GetRemoteManyMap(key string) (*map[string]Entity, error) {
  var entityResp ApricateResponse[map[string]Entity]
  _, err := HttpClientGet(&entityResp, service.UrlService.GetManyUrl(key))
  if err != nil {
    return nil, err
  }
  return &entityResp.Data, nil
}
