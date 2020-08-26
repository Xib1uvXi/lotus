package xcontext

import (
	"context"
	"time"
)

type XgP1Context struct {
	Ctx      context.Context
	HostName string
}

func NewXgP1Context(ctx context.Context) *XgP1Context {
	return &XgP1Context{Ctx: ctx}
}

func (x *XgP1Context) Deadline() (deadline time.Time, ok bool) {
	panic("implement me")
}

func (x *XgP1Context) Done() <-chan struct{} {
	panic("implement me")
}

func (x *XgP1Context) Err() error {
	panic("implement me")
}

func (x *XgP1Context) Value(key interface{}) interface{} {
	return x.HostName
}
