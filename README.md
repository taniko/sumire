# sumire
Sumire is logger library.

## Feature
- Log level are DEBUG, INFO, NOTICE, WARNING, ERROR, CRITICAL, ALERT, and EMERGENCY
- Record context value to log. (optional)
- Record runtime file and line to log. (optional)
- Support custom handler. (e.g., Send notification when called Alert)


## Guide
## Installation
```shell
go get github.com/taniko/sumire
```

### Example
```go
package main

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/taniko/sumire"
)

type requestContextKey struct{}

func main() {
	rand.Seed(time.Now().UnixNano())

	logger := sumire.NewLogger("sumire",
		sumire.WithStandardHandler(sumire.INFO, os.Stdout),
		sumire.WithRuntimeExtra(sumire.WARNING, sumire.DefaultRuntimeSkip),
		sumire.WithAppendixExtra("request-id", requestContextKey{}),
	)

	ctx := context.Background()
	// Recommend to use UUID for value
	ctx = context.WithValue(ctx, requestContextKey{}, rand.Int())

	// Skip output. Because handler has set INFO level.
	logger.DebugContext(ctx, "call debug", nil)

	// Output with request-id in extra
	logger.InfoContext(ctx, "call info", map[string]interface{}{
		"key": "value",
	})

	// Output with file and line in extra
	logger.WarningContext(ctx, "call warning", nil)
}
```