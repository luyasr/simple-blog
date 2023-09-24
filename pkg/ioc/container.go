package ioc

var (
	controllerContainer = &Container{
		store: map[string]Ioc{},
	}
	apiHandlerContainer = &Container{
		store: map[string]Ioc{},
	}
)

func Controller() *Container {
	return controllerContainer
}

func ApiHandler() *Container {
	return apiHandlerContainer
}
