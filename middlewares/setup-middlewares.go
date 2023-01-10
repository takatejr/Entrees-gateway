package middlewares

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func SetupMiddlewares(router *http.ServeMux) {
	for _, route := range ROUTES {
		fn := http.HandlerFunc(route.Controller)
		defaultHandler := RequestLogger(fn, route.Url)

		if route.Auth {
			router.Handle(route.Url, CheckToken(defaultHandler))
		} else {
			router.Handle(route.Url, defaultHandler)
		}

	}
	logrus.Println("Routes & Middlewares ready.")
}
