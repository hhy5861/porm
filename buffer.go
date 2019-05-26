package porm

import "strings"

type Buffer interface {
	WriteString(string) (int, error)
	String() string

	WriteValue(v ...interface{}) (err error)
	Value() []interface{}
}

type buffer struct {
	strings.Builder
	v []interface{}
}

func NewBuffer() Buffer {
	return &buffer{}
}

func (b *buffer) WriteValue(v ...interface{}) error {
	b.v = append(b.v, v...)
	return nil
}

func (b *buffer) Value() []interface{} {
	return b.v
}
