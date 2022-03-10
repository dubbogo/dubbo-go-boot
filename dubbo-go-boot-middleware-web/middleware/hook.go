package middleware

import (
	"github.com/gin-gonic/gin"
)

type WebSetupHook struct {
	hook func(router *gin.Engine)
}

func NewWebSetupHook(hook func(router *gin.Engine)) *WebSetupHook {
	return &WebSetupHook{
		hook: hook,
	}
}

func (m *WebSetupHook) Hook() {

}
