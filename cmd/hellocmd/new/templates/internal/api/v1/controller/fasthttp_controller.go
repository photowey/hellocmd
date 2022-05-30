package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"codeup.aliyun.com/uphicoo/gokit/expression"
	"codeup.aliyun.com/uphicoo/gokit/log"
	"codeup.aliyun.com/uphicoo/gokit/stopwatch"
	"codeup.aliyun.com/uphicoo/gokit/timez"
	perrors "github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"uphicoo.com/uphicoo/project-template/internal/api/router"
	"uphicoo.com/uphicoo/project-template/internal/config"
	"uphicoo.com/uphicoo/project-template/internal/pkg/tracer"
)

// Controller web http 请求处理器
type Controller struct {
	Host string
}

// NewController 构建 http web 请求处理器
func NewController(host string) *Controller {
	return &Controller{
		Host: host,
	}
}

// Run 运行 http web Controller
func (h *Controller) Run(cancelAble context.Context, crush chan error) {
	go func() {
		crush <- fasthttp.ListenAndServe(h.Host, h.HandleFastHTTP)
	}()

	go func() {
		if err := healthz(cancelAble); err != nil {
			log.Error(err.Error())
		} else {
			if log.IsInfoEnabled() {
				log.Infof("the server:[pid:%d] has been deployed on:[%s](HTTP) successfully.", os.Getpid(), config.Host())
			}
		}
	}()
}

func healthz(cancelAble context.Context) error {
	host := config.Host()
	url := fmt.Sprintf("http://%s/healthz", config.Host())

	bind := strings.Split(host, ":")[0]
	if bind == "" || bind == "0.0.0.0" {
		url = fmt.Sprintf("http://127.0.0.1:%s/healthz", strings.Split(host, ":")[1])
	}

	for {
		request, err := http.NewRequestWithContext(cancelAble, http.MethodGet, url, nil)
		if err != nil {
			return err
		}
		resp, err := http.DefaultClient.Do(request)
		if err == nil && resp.StatusCode == http.StatusOK {
			_ = resp.Body.Close()

			return nil
		}

		select {
		case <-cancelAble.Done():
			timeout := expression.TrinaryOperationUInt32(config.Timeout() > 0, config.Timeout(), 15)
			return perrors.Errorf("can not ping http server within the specified time:[%d s] interval", timeout)
		default:
		}
	}
}

// HandleFastHTTP 处理http 请求
func (h *Controller) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	handleStart := timez.Now()
	stopWatch := stopwatch.NewStopWatch(handleStart)

	path := string(ctx.Path())
	router.Route(ctx)

	report := formatReport("path:[%s], total", path)
	tracer.ReportRequest(report, stopWatch.ElapsedMicro())
}

// ---------------------------------------------------------------- 内部函数

// formatReport 格式化: 请求报告文案
func formatReport(message string, args ...any) string {
	return fmt.Sprintf(message, args...)
}
