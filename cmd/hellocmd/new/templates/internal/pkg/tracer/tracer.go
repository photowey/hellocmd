package tracer

import (
	"codeup.aliyun.com/uphicoo/gokit/log"
)

// ReportRequest 报告请求执行情况
func ReportRequest(message string, consume string) {
	log.Infof("tracer: %s cost: %s", message, consume)
}
