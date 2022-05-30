package router

import (
	"log"
	"sort"
	"sync"

	"codeup.aliyun.com/uphicoo/gokit/pkg/httpz"
	"github.com/valyala/fasthttp"

	"uphicoo.com/uphicoo/project-template/internal/api/handler"
)

var (
	_registry *Registry
	_handlers []handler.RequestHandler
	_lock     sync.RWMutex
)

func init() {
	_registry = newRegistry()
}

type Registry struct {
	ctx map[string]handler.RequestHandler
}

func (registry Registry) register(path string, handler handler.RequestHandler) {
	_lock.Lock()
	defer _lock.Unlock()
	if handler == nil {
		panic("handleregistry: web.http.handler.registry: Register handler is nil")
	}

	if _, dup := registry.ctx[path]; dup {
		panic("handleregistry: web.http.handler.registry: Register called twice for path: " + path)
	}

	registry.ctx[path] = handler
	log.Printf("http.router: register web.handler:[%s]", path)
}

func (registry Registry) handlers() []handler.RequestHandler {
	handlers := make([]handler.RequestHandler, 0)

	for _, handler := range registry.ctx {
		handlers = append(handlers, handler)
	}

	// handler 排序
	sort.Slice(handlers, func(i, j int) bool {
		return handlers[i].Order() < handlers[j].Order()
	})

	return handlers
}

func newRegistry() *Registry {
	return &Registry{
		ctx: make(map[string]handler.RequestHandler),
	}
}

// ----------------------------------------------------------------

// Register - 注册 handler
func Register(path string, handler handler.RequestHandler) {
	_registry.register(path, handler)
}

// CollectHandler - 收集 handler 列表
func CollectHandler() {
	if nil == _handlers || len(_handlers) == 0 {
		_lock.Lock()
		defer _lock.Unlock()
		if nil == _handlers || len(_handlers) == 0 {
			_handlers = _registry.handlers() // 为什么? 避免每次请求来的时候 - 再去构造 slice
		}
	}
}

// Route 根据请求地址 path 路由到合适的 Handler
func Route(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType(httpz.ContentTypeApplicationJSON)
	ctx.SetStatusCode(fasthttp.StatusOK)

	dispatch(ctx)
}

func dispatch(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	for _, handler := range _handlers {
		if handler.Supports(path) {
			handler.Handle(ctx)
			return
		}
	}
}
