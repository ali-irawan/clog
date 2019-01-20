# clog
A Centralized Logger for Golang. Inspired by `database/sql` package.

## Usage
```go
package main

import (
  "github.com/nbs-go/clog"
  // Register a logger implementation, just do it once in the main package
  _ "github.com/nbs-go/clogrus"
)

func main() {
  // Get logger singleton
  log := clog.Get()
  log.Info("Hello World")
}
```
