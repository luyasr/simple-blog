package ioc

import "github.com/gin-gonic/gin"

type Container struct {
	store map[string]Ioc
}

// Init 注册ioc容器中的对象
func (c *Container) Init() error {
	for _, obj := range c.store {
		if err := obj.Init(); err != nil {
			return nil
		}
	}
	return nil
}

func (c *Container) Registry(obj Ioc) {
	c.store[obj.Name()] = obj
}

func (c *Container) Get(name string) any {
	return c.store[name]
}

// RouteRegistry 注册路由到gin route
func (c *Container) RouteRegistry(r gin.IRouter) {
	for _, route := range c.store {
		if api, ok := route.(GinRouterGroup); ok {
			api.Registry(r)
		}
	}
}
