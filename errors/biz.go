package errors

import (
	"github.com/ascode/horde-sdk/stack"
	"google.golang.org/grpc/status"
)

type BizErr struct {
	Msg string
	fx  stack.Fx
}

func NewBizErr(msg string, kv ...interface{}) *BizErr {
	e := BizErr{
		Msg: msg,
		fx:  stack.Fx{},
	}

	if len(kv)%2 == 0 {
		for i := 0; i < len(kv); i += 2 {
			if k, ok := kv[i].(string); ok {
				e.fx[k] = kv[i+1]
			}
		}
	}

	return &e
}

func (e *BizErr) Error() string { return e.Msg }

func (e *BizErr) Fx() stack.Fx { return e.fx }

func (e *BizErr) GRPCStatus() *status.Status {
	return status.New(1118, e.Msg)
}

func IsBizErr(err error) (*BizErr, bool) {
	if e, ok := err.(*BizErr); ok {
		return e, true
	}
	return nil, false
}
