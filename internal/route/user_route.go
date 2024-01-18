package route

import (
	"github.com/gin-gonic/gin"
	"wallet/internal/handler"
	"wallet/internal/middleware"
)

func (r *Router) User(route *gin.RouterGroup, h *handler.Handler) {
	route.GET("/profiles", middleware.AuthMiddleware(r.jwtService, r.userService), h.Profile)
	route.POST("/cards", middleware.AuthMiddleware(r.jwtService, r.userService), h.Cards)
}
