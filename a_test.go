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

func Test_newObj(t *testing.T) {

}

func Test_obj_Bytes(t *testing.T) {

}

func Test_obj_String(t *testing.T) {

}
