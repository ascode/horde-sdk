package errors

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type Frame struct {
	Name string
}

func GetFrame(trackId string, skip int) *Frame {
	pc, file, ln, ok := runtime.Caller(1 + skip)
	if !ok {
		return nil
	}
	return &Frame{
		Name: fmt.Sprintf("trackId:%s,%s:%v:%v", trackId, file, filepath.Base(runtime.FuncForPC(pc).Name()), ln),
	}
}
