package controller

import (
	"codeup.aliyun.com/uphicoo/gokit/typez"
	"github.com/valyala/fasthttp"

	"uphicoo.com/uphicoo/project-template/internal/api/router"
)

func init() {
	router.Register(RequestPathDefault, &DefaultHandler{})
}

type DefaultHandler struct{}

// Order 序号
//
// 永远排在最后
func (h DefaultHandler) Order() int64 {
	return typez.Int64Max
}

func (h DefaultHandler) Supports(path string) bool {
	return true
}

func (h DefaultHandler) Handle(ctx *fasthttp.RequestCtx) {
	ctx.Error("web.http.dispatcher: unsupported path", fasthttp.StatusNotFound)
}
