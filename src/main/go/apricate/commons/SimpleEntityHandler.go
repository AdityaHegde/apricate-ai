package commons

type SimpleEntityHandler[Entity interface{}] struct {
	EntityHandler[Entity]
	LoadedHandler func(entity *Entity)
}

func (handler *SimpleEntityHandler[Entity]) EntityLoaded(entity *Entity) {
	handler.LoadedHandler(entity)
}
