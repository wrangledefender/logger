// Copyright 2021 wranglerdefender <wrangler.defender@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	li   *zap.Logger
	once sync.Once
)

func Configuration(ops ...Option) {
	options := &options{
		output:      []string{"stdout"},
		timeEncoder: zapcore.RFC3339TimeEncoder,
	}

	for _, o := range ops {
		o.apply(options)
	}

	li = NewLoggerOr(ops...)
}

// NewLoggerOr Initialize zap li instance.
func NewLoggerOr(ops ...Option) *zap.Logger {
	var err error
	var options = &options{
		output:      []string{"stdout"},
		timeEncoder: zapcore.RFC3339TimeEncoder,
		encoding:    "json",
	}

	for _, o := range ops {
		o.apply(options)
	}

	once.Do(func() {
		config := zap.NewProductionConfig()
		config.OutputPaths = options.output
		config.EncoderConfig.EncodeTime = options.timeEncoder
		config.Encoding = options.encoding

		li, err = config.Build(zap.AddCallerSkip(1))
		if err != nil {
			panic(err)
		}
	})

	return li
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	logger := NewLoggerOr()
	defer logger.Sync()
	logger.Sugar().Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	logger := NewLoggerOr()
	defer logger.Sync()
	logger.Sugar().Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	logger := NewLoggerOr()
	defer logger.Sync()
	logger.Sugar().Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	logger := NewLoggerOr()
	defer logger.Sync()
	logger.Sugar().Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	logger := NewLoggerOr()
	defer logger.Sync()
	logger.Sugar().DPanicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	logger := NewLoggerOr()
	defer logger.Sync()
	logger.Sugar().Fatalw(msg, keysAndValues...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	logger := NewLoggerOr()
	defer logger.Sync()
	logger.Sugar().Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	logger := NewLoggerOr()
	defer logger.Sync()
	logger.Sugar().Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	logger := NewLoggerOr()
	defer logger.Sync()
	logger.Sugar().Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	logger := NewLoggerOr()
	defer logger.Sync()
	logger.Sugar().Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	logger := NewLoggerOr()
	defer logger.Sync()
	logger.Sugar().DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	logger := NewLoggerOr()
	defer logger.Sync()
	logger.Sugar().Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	logger := NewLoggerOr()
	defer logger.Sync()
	logger.Sugar().Fatalf(template, args...)
}
