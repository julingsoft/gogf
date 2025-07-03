package plugins

import "github.com/gin-gonic/gin"

type Plugin interface {
	Name() string
	Init() error
	Destroy() error
	Router() RouterFunc
}

type RouterFunc func(*gin.Engine)
