package utils

import (
	"fmt"
	"runtime"
)

var (
	UserAgent = fmt.Sprintf(
		"%s_%s:goradiooooo (by /u/olshevskiy87)",
		runtime.GOOS, runtime.GOARCH,
	)
)
