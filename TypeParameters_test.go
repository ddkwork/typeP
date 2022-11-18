package main

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	d := New("22333")
	fmt.Println(d.MarshalBinary())
	d = New(88888)
	fmt.Println(d.MarshalBinary())
}
