package porm

import "context"

type EventReceiver interface {
	Event(eventName string)
	EventKv(eventName string, kvs map[string]string)
	EventErr(eventName string, err error) error
	EventErrKv(eventName string, err error, kvs map[string]string) error
	Timing(eventName string, nanoseconds int64)
	TimingKv(eventName string, nanoseconds int64, kvs map[string]string)
}

type TracingEventReceiver interface {
	SpanStart(ctx context.Context, eventName, query string) context.Context
	SpanError(ctx context.Context, err error)
	SpanFinish(ctx context.Context)
}

type kvs map[string]string

var nullReceiver = &NullEventReceiver{}

type NullEventReceiver struct{}

func (n *NullEventReceiver) Event(eventName string) {}

func (n *NullEventReceiver) EventKv(eventName string, kvs map[string]string) {}

func (n *NullEventReceiver) EventErr(eventName string, err error) error { return err }

func (n *NullEventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	return err
}

func (n *NullEventReceiver) Timing(eventName string, nanoseconds int64) {}

func (n *NullEventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {}
