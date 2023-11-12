package ioc

var (
	controllerContainer = &Container{
		store: map[string]Ioc{},
	}
	httpHandlerContainer = &Container{
		store: map[string]Ioc{},
	}
)

func Controller() *Container {
	return controllerContainer
}

func HttpHandler() *Container {
	return httpHandlerContainer
}
