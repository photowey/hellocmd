package controller

import (
	"codeup.aliyun.com/uphicoo/gokit/jsonz"
	"github.com/valyala/fasthttp"

	"uphicoo.com/uphicoo/project-template/internal/api/router"
	"uphicoo.com/uphicoo/project-template/internal/types"
)

func init() {
	router.Register(RequestPathHealthz, &HealthHandler{})
}

type HealthHandler struct{}

func (h HealthHandler) Order() int64 {
	return 200
}

func (h HealthHandler) Supports(path string) bool {
	return path == RequestPathHealthz
}

// Handle
//
// healthz GET /healthz
func (h HealthHandler) Handle(ctx *fasthttp.RequestCtx) {
	up := types.NewUpResponse()
	bodyz := jsonz.String(up)
	ctx.SetBody([]byte(bodyz))
}
