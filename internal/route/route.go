package route

import "wallet/internal/service"

type Router struct {
	userService service.UserService
	jwtService  service.JWTService
}

type RouterConfig struct {
	UserService service.UserService
	JWTService  service.JWTService
}

func NewRouter(c *RouterConfig) *Router {
	return &Router{
		userService: c.UserService,
		jwtService:  c.JWTService,
	}
}
