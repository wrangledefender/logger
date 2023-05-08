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

import "go.uber.org/zap/zapcore"

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (o optionFunc) apply(ops *options) {
	o(ops)
}

type options struct {
	output      []string
	timeEncoder zapcore.TimeEncoder
	encoding    string
}

// WithEncoding Specify logger format, the valid "json" and "console".
func WithEncoding(encoding string) Option {
	return optionFunc(func(ops *options) {
		ops.encoding = encoding
	})
}

func WithOutput(output ...string) Option {
	return optionFunc(func(ops *options) {
		ops.output = output
	})
}

func WithTimeEncoder(timeEncoder zapcore.TimeEncoder) Option {
	return optionFunc(func(ops *options) {
		ops.timeEncoder = timeEncoder
	})
}
