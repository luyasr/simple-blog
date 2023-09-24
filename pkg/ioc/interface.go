package ioc

import "github.com/gin-gonic/gin"

type Ioc interface {
	Init() error
	Name() string
}

type GinRouterGroup interface {
	Registry(r gin.IRouter)
}
