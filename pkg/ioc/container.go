package ioc

var (
	controllerContainer = &Container{
		store: map[string]Ioc{},
	}
	handlerContainer = &Container{
		store: map[string]Ioc{},
	}
)

func Controller() *Container {
	return controllerContainer
}

func Handler() *Container {
	return handlerContainer
}
