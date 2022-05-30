package types

import (
	"strings"

	"codeup.aliyun.com/uphicoo/gokit/pkg/nanoid"
	"github.com/valyala/fasthttp"
)

const (
	RequestIdLength int = 21 // 请求参数标识-长度
)

// ----------------------------------------------------------------

// IRequest 请求-接口抽象
type IRequest interface {
	InjectId(requestId string)
	InjectPath(path string)
	InjectMethod(method string)
	InjectCtx(ctx *fasthttp.RequestCtx)
}

func requestInjected(ctx *fasthttp.RequestCtx, request IRequest) {
	requestId, _ := nanoid.New(RequestIdLength)
	request.InjectId(requestId)
	request.InjectPath(parseRequestPath(ctx))
	request.InjectMethod(parseHttpMethod(ctx))
	request.InjectCtx(ctx)
}

// Request 请求数据模型
type Request struct {
	Id     string               `json:"id"`     //  请求标识
	Path   string               `json:"path"`   // 请求地址
	Method string               `json:"method"` // http 请求 Method
	Ctx    *fasthttp.RequestCtx `json:"ctx"`    //  请求 上下文
}

func (r *Request) InjectId(requestId string) {
	r.Id = requestId
}

func (r *Request) InjectPath(path string) {
	r.Path = path
}

func (r *Request) InjectMethod(method string) {
	r.Method = method
}

func (r *Request) InjectCtx(ctx *fasthttp.RequestCtx) {
	r.Ctx = ctx
}

// ----------------------------------------------------------------

type HealthResponse struct {
	Status string `json:"status"`
}

func NewUpResponse() HealthResponse {
	return HealthResponse{
		Status: "UP",
	}
}

func NewDownResponse() HealthResponse {
	return HealthResponse{
		Status: "DOWN",
	}
}

// IResponse 请求响应-接口抽象
type IResponse interface {
	InjectCode(code string)
	InjectMessage(message string)
}

// ----------------------------------------------------------------

// Response 响应数据模型基类
type Response struct {
	Code    string `json:"code"`    //  请求返回嘛
	Message string `json:"message"` // 请求消息
}

func (r *Response) InjectCode(code string) {
	r.Code = code
}

func (r *Response) InjectMessage(message string) {
	r.Message = message
}

func parseHttpMethod(ctx *fasthttp.RequestCtx) string {
	method := string(ctx.Method())

	return strings.ToUpper(method)
}

func parseRequestPath(ctx *fasthttp.RequestCtx) string {
	path := string(ctx.Path())

	return path
}
