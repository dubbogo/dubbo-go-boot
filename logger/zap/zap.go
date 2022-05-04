/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package zap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dubbogo/dubbo-go-boot/core"
	"github.com/dubbogo/dubbo-go-boot/core/constant"
	"github.com/dubbogo/dubbo-go-boot/core/extension"
	"github.com/dubbogo/dubbo-go-boot/logger"
)

func init() {
	extension.SetLogger("zap", newZapLogger)
}

type Logger struct {
	lg *zap.SugaredLogger
}

func newZapLogger(conf *core.URL) (log logger.Logger, err error) {
	level := conf.GetParam(constant.LoggerLevelKey, "info")
	if log, err = getLogger(level); err != nil {
		return nil, err
	}
	return log, nil
}

func (l *Logger) Debug(args ...interface{}) {
	l.lg.Debug(args)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.lg.Debugf(template, args)
}

func (l *Logger) Info(args ...interface{}) {
	l.lg.Info(args)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.lg.Infof(template, args)
}

func (l *Logger) Warn(args ...interface{}) {
	l.lg.Warn(args)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.lg.Warnf(template, args)
}

func (l *Logger) Error(args ...interface{}) {
	l.lg.Error(args)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.lg.Errorf(template, args)
}

func getEncoder() zapcore.Encoder {
	encoder := zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			CallerKey:      "line",
			NameKey:        "logger",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
	return encoder
}

func getLogger(level string) (*zap.SugaredLogger, error) {
	var (
		lv  zapcore.Level
		err error
	)
	if lv, err = zapcore.ParseLevel(level); err != nil {
		return nil, err
	}
	encoder := getEncoder()
	return zap.New(zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lv),
		zap.AddCaller(), zap.AddCallerSkip(1)).Sugar(), nil
}
