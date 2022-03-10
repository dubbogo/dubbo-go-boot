package component

import (
	"github.com/gin-gonic/gin"
)

var (
	WebComponent = &webComponent{}
)

type webComponent struct {
	Router *gin.Engine
}

func GetRouter() *gin.Engine {
	return WebComponent.Router
}
