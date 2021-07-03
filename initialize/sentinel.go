package initialize

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
)

func InitSentinel() error {
	err := sentinel.InitDefault()
	if err != nil {
		return err
	}

	// 配置限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "GET:/api/v1/lecturers",
			TokenCalculateStrategy: flow.Direct, // Direct: 直接拒绝, Throttling: 匀速通过
			ControlBehavior:        flow.Reject, // Reject: qps限流, WarmUp: 预热方式
			Threshold:              10,
			StatIntervalInMs:       1000, // 1秒10个并发
		},
	})

	if err != nil {
		return err
	}

	return nil
}
