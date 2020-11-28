package middlewares

import "net/http"

func TokenAuthMiddleware() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
