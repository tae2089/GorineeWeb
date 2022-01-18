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
	SetCookie(key, value string, expires time.Duration, httpOnly, secure bool) bool
	GetCookie(key string) string
	Reply(contentType string, data []byte)
}

type Context interface {
	Next()
	Context() *fasthttp.RequestCtx
	Param(key string) string
	Query(key string) string
	SendBytes(value []byte) Context
	SendString(value string) Context
	SendJSON(in interface{}) error
	Status(status int) Context
	Set(key string, value string)
	Get(key string) string
	SetLocal(key string, value interface{})
	GetLocal(key string) interface{}
	Body() string
	ParseBody(out interface{}) error
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

//func (g *gorineeWebContext) Reply
