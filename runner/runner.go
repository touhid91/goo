package runner

import (
	"os"
	"time"
)

type Runner struct {
	interrupt chan os.Signal // represents signal from os channel reports
	complete chan error // reports processing done
	timeout <- chan time.Time
}
