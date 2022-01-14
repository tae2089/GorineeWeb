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

type GorineeWeb interface {
	Get(path string, handlers fasthttp.RequestHandler) *router.Router
	Post()
	New() GorineeWeb
	Run(addr string)
}

func (g *gorineeWeb) New() GorineeWeb {
	return &gorineeWeb{
		httpServer: &fasthttp.Server{},
		router:     router.New(),
		settings:   &Settings{},
	}
}

func (g *gorineeWeb) Run(addr string) {
	g.httpServer.Handler = g.router.Handler
	g.httpServer.ListenAndServe(addr)
}
func (g *gorineeWeb) Get(path string, handlers fasthttp.RequestHandler) *router.Router {
	//TODO implement me
	panic("implement me")
}

func (g *gorineeWeb) Post() {
	//TODO implement me
	panic("implement me")
}
