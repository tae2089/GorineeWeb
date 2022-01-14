package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

//router 생성하기

type gorineeWeb struct {
	httpServer *fasthttp.Server
	router     *router.Router
	//registeredRoutes []*Route
	address string // server address
	//middlewares      handlersChain
	settings *Settings
}

func (g gorineeWeb) Get(path string, handlers fasthttp.RequestHandler) *Router {
	//TODO implement me
	panic("implement me")
}

func (g gorineeWeb) Post() {
	//TODO implement me
	panic("implement me")
}

func (g gorineeWeb) New() GorineeWeb {
	//TODO implement me
	panic("implement me")
}

type Router struct {
}

type GorineeWeb interface {
	Get(path string, handlers fasthttp.RequestHandler) *Router
	Post()
	New() GorineeWeb
}

func New(settings ...*Settings) GorineeWeb {
	gorin := new(gorineeWeb)

	if len(settings) > 0 {
		gorin.settings = settings[0]
	} else {
		gorin.settings = &Settings{}
	}

	//if gorin.settings.ServerName <= 0 {
	//	gorin.settings.ServerName =
	//}
	return gorin
}

func DefaultRouter() *router.Router {
	r := router.New()
	return r
}
