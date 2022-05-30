package handler

import (
	"codeup.aliyun.com/uphicoo/gokit/orderz"
	"github.com/valyala/fasthttp"
)

// RequestHandler - fasthttp request handler
type RequestHandler interface {
	orderz.Ordered
	Supports(path string) bool
	Handle(ctx *fasthttp.RequestCtx)
}
