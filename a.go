package main

import "fmt"

type (
	T interface {
		string | []byte
	}
	IN[t T] interface {
		fmt.Stringer
		Bytes() []byte
	}
	obj[t T] struct {
		data t
	}
)

func (o *obj[t]) Bytes() []byte {
	return o.data.([]byte)
}

func (o *obj[t]) String() string {
	return o.data.(string)
}

func newObj[t T](data t) *obj[t] {
	return &obj[t]{data: data}
}

func New[t T](data t) IN[t] {
	return newObj(data)
}
