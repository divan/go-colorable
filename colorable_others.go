// +build !windows

package colorable

import (
	"io"
	"os"

	_ "github.com/mattn/go-isatty"
)

func NewColorableStdout() io.Writer {
	return os.Stdout
}

func NewColorableStderr() io.Writer {
	return os.Stderr
}
