package ioc

import "github.com/gin-gonic/gin"

type Ioc interface {
	Init() error
	Name() string
}

type GinIRouter interface {
	Registry(r gin.IRouter)
}
