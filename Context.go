package GorineeWeb

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
type GorineeWebContext struct {
	ctx *fasthttp.RequestCtx
}

type gorineeWebContext interface {
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

func (g *GorineeWebContext) SetCookie(key, value string, expireAt time.Time, httpOnly, secure bool) bool {
	ck := fasthttp.AcquireCookie()
	ck.SetKey(key)
	ck.SetValue(value)
	ck.SetHTTPOnly(httpOnly)
	ck.SetSecure(secure)
	ck.SetExpire(expireAt)
	success := g.ctx.Response.Header.Cookie(ck)
	return success
}

func (g *GorineeWebContext) GetCookie(key string) string {
	return string(g.ctx.Request.Header.Cookie(key))
}

func (g *GorineeWebContext) Reply(contentType string, data []byte) {
	g.ctx.SetContentType(contentType)
	g.ctx.SetStatusCode(fasthttp.StatusOK)
	g.ctx.Response.Header.Set("Content-Length", strconv.FormatInt(int64(len(data)), 10))
	g.ctx.Response.SetBody(data)
}

func (g *GorineeWebContext) ReplyJSON(data interface{}) {
	g.ctx.SetContentType("application/json")
	g.ctx.SetStatusCode(fasthttp.StatusOK)
	encoder := json.NewEncoder(g.ctx.Response.BodyWriter())
	if err := encoder.Encode(data); err != nil {
		//g.Log.WriteLog(logger.SYSTEM, log.New(loglevel.Error, err.Error()).End())
		g.ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
}

func (g *GorineeWebContext) Status(status STATUS) {
	g.ctx.Response.SetStatusCode(status.convert())
}

func (g *GorineeWebContext) Param(key string) interface{} {
	return g.ctx.UserValue(key)
}

func (g *GorineeWebContext) Query(key string) string {
	return GetString(g.ctx.QueryArgs().Peek(key))
}

func (g *GorineeWebContext) Body() string {
	return GetString(g.ctx.Request.Body())
}
