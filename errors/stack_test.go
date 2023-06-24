package errors

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestTraceError(t *testing.T) {
	err := errfunc3()
	if e, ok := err.(*StackErr); ok {
		cause := e.Cause()
		if be, ok := IsBizErr(cause); ok {
			log.Error(cause.Error(), e.FxStack(), be.fx)
		} else {
			log.Error(cause.Error(), e.FxStack())
		}
	}
}

func errfunc3() error {
	err := errfunc2()
	return WrapEx("id0", 1, err)
}

func errfunc2() error {
	err := errfunc1()
	return WrapEx("id0", 1, err)
}

func errfunc1() error {
	err := NewBizErr("deep err", "name", "9527", "size", 18)
	return WrapEx("id0", 1, err)
}
