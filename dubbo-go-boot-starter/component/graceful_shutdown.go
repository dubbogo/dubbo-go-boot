package component

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/config"
	"os"
	"os/signal"
	"time"
)

func ObserveSignal(duration time.Duration, beforeShutdown func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, config.ShutdownSignals...)

	for {
		select {
		case sig := <-signals:
			logger.Infof("get signal %s, applicationConfig will shutdown.", sig)
			// gracefulShutdownOnce.Do(func() {
			time.AfterFunc(duration, func() {
				logger.Warn("Shutdown gracefully timeout, applicationConfig will shutdown immediately. ")
				os.Exit(0)
			})

			if beforeShutdown != nil {
				beforeShutdown()
			}
			// those signals' original behavior is exit with dump ths stack, so we try to keep the behavior
			os.Exit(0)
		}
	}
}
