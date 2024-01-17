package route

import s "wallet/internal/service"

type Router struct {
	userService s.UserService
	jwtService  s.JWTService
}

type RouterConfig struct {
	UserService s.UserService
	JWTService  s.JWTService
}

func NewRouter(c *RouterConfig) *Router {
	return &Router{
		userService: c.UserService,
		jwtService:  c.JWTService,
	}
}
