package basic

import (
	"fmt"
	"testing"
	"io"
	"reflect"
)

type decoder interface {
	decode(io.Reader)
}

type decoderImpl struct {
	size uint64
}

func (*decoderImpl) decode(io.Reader) {

}

type decoderImpl1 struct {
	size uint64
}

func (decoderImpl1) decode(io.Reader) {

}

type decoderWrap decoderImpl

var decoderInterface = reflect.TypeOf(new(decoder)).Elem()

func TestReflect0(t *testing.T) {
	di := new(decoderImpl)
	typ := reflect.TypeOf(di)
	val := reflect.ValueOf(di)
	fmt.Println(val.Kind()) // ptr
	fmt.Println(val.IsNil()) // false
	fmt.Println(typ.Implements(decoderInterface)) // true

	var di1 *decoderImpl
	typ1 := reflect.TypeOf(di1)
	val1 := reflect.ValueOf(di1)
	fmt.Println(val1.Kind()) // ptr
	fmt.Println(val1.IsNil()) // true
	fmt.Println(typ1.Implements(decoderInterface)) // true

	di2 := decoderImpl{size:1}
	typ2 := reflect.TypeOf(di2)
	val2 := reflect.ValueOf(di2)
	fmt.Println(val2.Kind()) // struct
	//fmt.Println(val2.IsNil()) panic
	fmt.Println(typ2.Implements(decoderInterface)) // false

	dw := new(decoderWrap)
	typ3 := reflect.TypeOf(dw)
	fmt.Println(typ3.Implements(decoderInterface)) // false

	di3 := new(decoderImpl1)
	typ4 := reflect.TypeOf(di3)
	val4 := reflect.ValueOf(di3)
	fmt.Println(val4.Kind()) // ptr
	fmt.Println(val4.IsNil()) // false
	fmt.Println(typ4.Implements(decoderInterface)) // true

	di5 := decoderImpl1{1}
	typ5 := reflect.TypeOf(di5)
	val5 := reflect.ValueOf(di5)
	fmt.Println(val5.Kind()) // struct
	//fmt.Println(val5.IsNil()) panic
	fmt.Println(typ5.Implements(decoderInterface)) // true
}

