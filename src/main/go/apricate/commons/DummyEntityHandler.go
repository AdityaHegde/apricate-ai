package commons

type DummyEntityHandler[Entity interface{}] struct {
  EntityHandler[Entity]
}

func (handler *DummyEntityHandler[Entity]) EntityLoaded(entity *Entity) {
  // NOOP
}
