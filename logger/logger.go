/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logger

import "github.com/dubbogo/dubbo-go-boot/logger/zap"

type Logger interface {
	Debug(args ...interface{})

	Debugf(template string, args ...interface{})

	Info(args ...interface{})

	Infof(template string, args ...interface{})

	Warn(args ...interface{})

	Warnf(template string, args ...interface{})

	Error(args ...interface{})

	Errorf(template string, args ...interface{})
}

var logger Logger

func SetLog(log Logger) {
	logger = log
}

func SetDefaultLog() {
	log, _ := zap.GetLogger("info")
	SetLog(log)
}

// Info is info level
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Warn is warning level
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Error is error level
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Debug is debug level
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Infof is format info level
func Infof(fmt string, args ...interface{}) {
	logger.Infof(fmt, args...)
}

// Warnf is format warning level
func Warnf(fmt string, args ...interface{}) {
	logger.Warnf(fmt, args...)
}

// Errorf is format error level
func Errorf(fmt string, args ...interface{}) {
	logger.Errorf(fmt, args...)
}

// Debugf is format debug level
func Debugf(fmt string, args ...interface{}) {
	logger.Debugf(fmt, args...)
}
