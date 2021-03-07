package chi

import (
	"github.com/go-chi/chi"
	"github.com/valonekowd/clean-architecture/infrastructure/router"
)

type chiRouter struct {
	chi.Router
}

func NewRouter() router.Router {
	return &chiRouter{
		Router: chi.NewRouter(),
	}
}

func (c *chiRouter) Route(pattern string, fn func(r router.Router)) router.Router {
	return &chiRouter{
		Router: c.Router.Route(pattern, func(r chi.Router) {
			fn(&chiRouter{
				Router: r,
			})
		}),
	}
}
