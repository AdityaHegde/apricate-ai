package commons

type SimpleEntityFieldsService[Entity interface{}] struct {
  EntityFieldsService[Entity]
  PrimaryKey         string
  PrimaryValueGetter func(entity *Entity) string
}

func (service *SimpleEntityFieldsService[Entity]) GetPrimaryKey() string {
  return service.PrimaryKey
}
func (service *SimpleEntityFieldsService[Entity]) GetPrimaryValue(entity *Entity) string {
  return service.PrimaryValueGetter(entity)
}
