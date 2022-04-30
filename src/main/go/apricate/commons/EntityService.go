package commons

import "gorm.io/gorm"

type EntityService[Entity interface{}] struct {
  Db            *gorm.DB
  FieldsService EntityFieldsService[Entity]
  Handler       EntityHandler[Entity]
  HttpService   HttpService[Entity]
}
type EntityFieldsService[Entity interface{}] interface {
  GetPrimaryKey() string
  GetPrimaryValue(entity *Entity) string
}
type EntityHandler[Entity interface{}] interface {
  EntityLoaded(entity *Entity)
}

// CreateOrSave either creates a fresh record or updates existing
func (service *EntityService[Entity]) CreateOrSave(entity *Entity) {
  existing := service.Db.First(
    entity,
    service.FieldsService.GetPrimaryKey()+" = ?",
    service.FieldsService.GetPrimaryValue(entity),
  )
  if existing.RowsAffected > 0 {
    service.Db.Save(entity)
    return
  }

  service.Db.Create(entity)
}

// GetOrFetch either gets entity from DB or fetches from remote
func (service *EntityService[Entity]) GetOrFetch(key string) (*Entity, error) {
  var entity Entity
  existing := service.Db.First(
    &entity,
    service.FieldsService.GetPrimaryKey()+" = ?",
    key,
  )
  if existing.RowsAffected > 0 {
    return &entity, nil
  }

  respEntity, err := service.HttpService.GetRemoteSingle(key)
  if err != nil {
    return nil, err
  }
  service.saveWrapper(respEntity)
  return respEntity, nil
}

func (service *EntityService[Entity]) LoadFromRemote(key string) (*Entity, error) {
  respEntity, err := service.HttpService.GetRemoteSingle(key)
  if err != nil {
    return nil, err
  }
  service.saveWrapper(respEntity)
  return respEntity, nil
}

func (service *EntityService[Entity]) LoadAllFromRemote(key string) (*[]Entity, error) {
  respEntities, err := service.HttpService.GetRemoteMany(key)
  if err != nil {
    return nil, err
  }
  for _, entity := range *respEntities {
    service.saveWrapper(&entity)
  }
  return respEntities, nil
}

func (service *EntityService[Entity]) LoadAllMapFromRemote(key string) (*map[string]Entity, error) {
  respEntities, err := service.HttpService.GetRemoteManyMap(key)
  if err != nil {
    return nil, err
  }
  for _, entity := range *respEntities {
    service.saveWrapper(&entity)
  }
  return respEntities, nil
}

func (service *EntityService[Entity]) saveWrapper(entity *Entity) {
  if service.FieldsService.GetPrimaryValue(entity) == "" {
    return
  }
  service.Handler.EntityLoaded(entity)
  service.Db.Save(entity)
}
