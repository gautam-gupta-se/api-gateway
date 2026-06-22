package router

import (
	"api-gateway/internal/model"
	"strings"
)

type Router struct {
	routes []model.Route
}

func NewRouter(routes []model.Route) *Router {
	return &Router{routes: routes}
}

func (r *Router) Match(path string) *model.Route {
	for _, route := range r.routes {
		// Matches prefix (e.g., /users/123 matches /users)
		if strings.HasPrefix(path, route.Path) {
			return &route
		}
	}
	return nil
}
