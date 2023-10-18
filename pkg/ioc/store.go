package ioc

import (
	"github.com/gin-gonic/gin"
)

type Container struct {
	store map[string]Ioc
}

// Init 注册ioc容器中的对象
func (c *Container) Init() error {
	for _, obj := range c.store {
		if err := obj.Init(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Container) Get(name string) any {
	return c.store[name]
}

func (c *Container) Registry(obj Ioc) {
	if _, exists := c.store[obj.Name()]; !exists {
		c.store[obj.Name()] = obj
	}
}

// RouteRegistry 注册路由到gin route
func (c *Container) RouteRegistry(r gin.IRouter) {
	for _, route := range c.store {
		if api, ok := route.(GinIRouter); ok {
			api.Registry(r)
		}
	}
}
