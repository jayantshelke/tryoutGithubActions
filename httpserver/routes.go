package httpserver

import (
	"ProjectIdeas/monolith/api"
	"context"
	"log"
	"net/http"
)

type RouteMethods map[string]http.Handler
type Router struct {
	routes map[string]RouteMethods
}

type RouterOptions struct {
	Api api.APIer
}

func (r Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	log.Println("url path :", request.URL.Path)

	if _, ok := r.routes[request.URL.Path]; !ok {
		log.Println("cannot find the resource requested")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	h := r.routes[request.URL.Path]
	var handler http.Handler

	// this is going to panic since we don't define all methods
	// on each of the endpoints
	//TODO : handle panic for undefined method handlers
	switch request.Method {
	case http.MethodGet:
		handler = h[http.MethodGet]
	case http.MethodDelete:
		handler = h[http.MethodDelete]
	case http.MethodPost:
		handler = h[http.MethodPost]
	case http.MethodPut:
		handler = h[http.MethodPut]
	case http.MethodOptions:
		handler = h[http.MethodOptions]
	default:
		handler = h[http.MethodHead]
	}

	handler.ServeHTTP(writer, request)
	return

}

func NewRouter(ctx context.Context, options RouterOptions) Router {

	// define all routes here
	// along with their handlers

	userGetHandle := newUserGetHandler(ctx, options.Api)
	userDeleteHandle := newUserDeleteHandler(ctx, options.Api)

	var routes = make(map[string]RouteMethods)

	routes[`/v1/users`] = RouteMethods{http.MethodGet: userGetHandle, http.MethodHead: userGetHandle, http.MethodDelete: userDeleteHandle}

	return Router{
		routes: routes,
	}
}
