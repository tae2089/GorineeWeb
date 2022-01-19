package main

import (
	"github.com/fasthttp/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"log"
)

// Exported constants
const (
	Version = "0.0.1 alpha" // Version of GorineeWeb
	Name    = "GorineeWeb"  // Name of GorineeWeb
	banner2 = `
  ____                  _               __        __   _     
 / ___| ___   ___  _ __(_)_ __   ___  __\ \      / /__| |__  
| |  _ / _ \ / _ \| '__| | '_ \ / _ \/ _ \ \ /\ / / _ \ '_ \ 
| |_| | (_) | (_) | |  | | | | |  __/  __/\ V  V /  __/ |_) |
 \____|\___/ \___/|_|  |_|_| |_|\___|\___| \_/\_/ \___|_.__/  v%s
Listening on %s`
)

//router 생성하기

type gorineeWeb struct {
	httpServer       *fasthttp.Server
	router           *router.Router
	registeredRoutes []*Route
	address          string // server address
	middlewares      handlersChain
	settings         *Settings
}

type Route struct {
	Method  METHOD
	Path    string
	Handler Handler
}

type GorineeWeb interface {
	Get(path string, handler Handler)
	Post(path string, handler Handler)
	Put(path string, handler Handler)
	Delete(path string, handler Handler)
	Head(path string, handler Handler)
	New() GorineeWeb
	Run(addr string)
	HealthCheck()
	handle(method METHOD, path string, handler Handler)
	//WebSocket(path string, handler Handler)
}

func (g *gorineeWeb) New() GorineeWeb {
	return &gorineeWeb{
		httpServer: &fasthttp.Server{Logger: &customLogger{}},
		router:     router.New(),
		settings:   &Settings{},
	}
}

//서버 실행하기
func (g *gorineeWeb) Run(addr string) {
	g.httpServer.Handler = g.router.Handler
	log.Printf(banner2, Version, addr)
	g.httpServer.ListenAndServe(addr)
}

func (g *gorineeWeb) HealthCheck() {
	fasthandler := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	g.router.Handle(Get.convert(), "metrics", fasthandler)
}

func (g *gorineeWeb) handle(method METHOD, path string, handler Handler) {
	g.router.Handle(method.convert(), path, func(ctx *fasthttp.RequestCtx) {
		gorineeWebContext := &gorineeWebContext{ctx: ctx}
		handler(gorineeWebContext)
	})
}

//func (g *gorineeWeb) WebSocket(path string, handler Handler) {
//	g.router.Handle(method.convert(), path, func(ctx *fasthttp.RequestCtx) {
//		gorineeWebContext := &gorineeWebContext{ctx: ctx}
//		handler(gorineeWebContext)
//	})
//}

//func (g *gorineeWeb) registeredRoute(method METHOD, path string, handler Handler) *Route {
//	route := &Route{
//		Path:    path,
//		Method:  method,
//		Handler: handler,
//	}
//	g.registeredRoutes = append(g.registeredRoutes, route)
//	return route
//}

func (g *gorineeWeb) Get(path string, handler Handler) {
	//TODO implement me
	g.handle(Get, path, handler)
}

func (g *gorineeWeb) Post(path string, handler Handler) {
	//TODO implement me
	g.handle(Post, path, handler)
}

func (g *gorineeWeb) Put(path string, handler Handler) {
	//TODO implement me
	g.handle(Put, path, handler)
}

func (g *gorineeWeb) Delete(path string, handler Handler) {
	//TODO implement me
	g.handle(Delete, path, handler)
}

func (g *gorineeWeb) Head(path string, handler Handler) {
	//TODO implement me
	g.handle(Head, path, handler)
}

func registerMiddleware(handler Handler) {

}
