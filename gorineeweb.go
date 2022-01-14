package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
)

// Exported constants
const (
	Version = "0.0.1 alpha" // Version of gearbox
	Name    = "GorineeWeb"  // Name of gearbox
	// http://patorjk.com/software/taag/#p=display&f=Big%20Money-ne&t=Gearbox
	banner = `
  _____                     _                 
 / ____|                   | |                
| |  __   ___   __ _  _ __ | |__    ___ __  __
| | |_ | / _ \ / _' || '__|| '_ \  / _ \\ \/ /
| |__| ||  __/| (_| || |   | |_) || (_) |>  < 
 \_____| \___| \__,_||_|   |_.__/  \___//_/\_\ v%s
Listening on %s`

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
	log.Printf(banner2, Version, addr)
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
