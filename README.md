# logger

Uber zap logger wrapper impl.

## Installation

```shell script
$ go get -u github.com/wranglerdefender/logger
```


## How to use？

It's very simple to use, just out of the box, and you can view the documentation through [GoDoc](https://pkg.go.dev/github.com/wranglerdefender/logger).


## Example 

```go
package main

import (
	"github.com/wranglerdefender/logger"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger.Configuration(
		logger.WithOutput("app.log"), 
		logger.WithTimeEncoder(zapcore.RFC3339NanoTimeEncoder),
	)

	logger.Infof("upgrade status: %v", true)
	logger.Infow("upgrade info", "status", true)
}
```
