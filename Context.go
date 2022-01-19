package main

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
)

// handlerFunc defines the handler used by middleware as return value.
//type handlerFunc func(ctx Context)

// handlersChain defines a handlerFunc array.
//type handlersChain []handlerFunc
// Context interface
type gorineeWebContext struct {
	ctx *fasthttp.RequestCtx
}

type GorineeWebContext interface {
	SetCookie(key, value string, expireAt time.Time, httpOnly, secure bool) bool
	GetCookie(key string) string
	Body() string
	Query(key string) string
	Param(key string) interface{}
	Reply(contentType string, data []byte)
	ReplyJSON(data interface{})
}

// handlersChain defines a handlerFunc array.
type handlersChain []fasthttp.RequestHandler

func (g *gorineeWebContext) SetCookie(key, value string, expireAt time.Time, httpOnly, secure bool) bool {
	ck := fasthttp.AcquireCookie()
	ck.SetKey(key)
	ck.SetValue(value)
	ck.SetHTTPOnly(httpOnly)
	ck.SetSecure(secure)
	ck.SetExpire(expireAt)
	success := g.ctx.Response.Header.Cookie(ck)
	return success
}

func (g *gorineeWebContext) GetCookie(key string) string {
	return string(g.ctx.Request.Header.Cookie(key))
}

func (g *gorineeWebContext) Reply(contentType string, data []byte) {
	g.ctx.SetContentType(contentType)
	g.ctx.SetStatusCode(fasthttp.StatusOK)
	g.ctx.Response.Header.Set("Content-Length", strconv.FormatInt(int64(len(data)), 10))
	g.ctx.Response.SetBody(data)
}

func (g *gorineeWebContext) ReplyJSON(data interface{}) {
	g.ctx.SetContentType("application/json")
	g.ctx.SetStatusCode(fasthttp.StatusOK)
	encoder := json.NewEncoder(g.ctx.Response.BodyWriter())
	if err := encoder.Encode(data); err != nil {
		//g.Log.WriteLog(logger.SYSTEM, log.New(loglevel.Error, err.Error()).End())
		g.ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
}

func (g *gorineeWebContext) Status(status STATUS) GorineeWebContext {
	g.ctx.Response.SetStatusCode(status.convert())
	return g
}

func (g *gorineeWebContext) Param(key string) interface{} {
	return g.ctx.UserValue(key)
}

func (g *gorineeWebContext) Query(key string) string {
	return GetString(g.ctx.QueryArgs().Peek(key))
}

func (g *gorineeWebContext) Body() string {
	return GetString(g.ctx.Request.Body())
}
