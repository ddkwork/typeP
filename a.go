package main

import (
	"encoding"
)

type (
	T interface {
		string | []byte
	}
	Interface[t T] interface {
		encoding.BinaryMarshaler
		encoding.BinaryUnmarshaler
	}
	object[t T] struct{ safeValue any }
)

func New[t T](safeValue t) Interface[t] { return &object[t]{safeValue: safeValue} }
func (o *object[t]) MarshalBinary() (data []byte, err error) {
	switch o.safeValue.(type) {
	case string:
		return []byte(o.safeValue.(string)), nil
	default:
		return nil, err
	}
}

func (o *object[t]) UnmarshalBinary(data []byte) error {
	//TODO implement me
	panic("implement me")
}
