package type_assertion

import (
	"testing"
)

type myint int64

type Inccer interface {
	inc()
}

func (i *myint) inc() {
	*i = *i + 1
}

func Benchmark_NormalSwitch(b *testing.B) {
	i := new(myint)
	incNormalSwitch(i, b.N)
}

func incNormalSwitch(i *myint, n int) {
	for k := 0; k < n; k++ {
		i.inc()
	}
}

func Benchmark_InterfaceSwitch(b *testing.B) {
	i := new(myint)
	incInterfaceSwitch(i, b.N)
}
func incInterfaceSwitch(any Inccer, n int) {
	for k := 0; k < n; k++ {
		any.inc()
	}
}

func Benchmark_interfaceSwitch(b *testing.B) {
	i := new(myint)
	incnInterfaceSwitch(i, b.N)
}

func incnInterfaceSwitch(any interface{}, n int) {
	for k := 0; k < n; k++ {
		any.(Inccer).inc()
	}
}

func Benchmark_TypeSwitch(b *testing.B) {
	i := new(myint)
	incTypeSwitch(i, b.N)
}

func incTypeSwitch(any interface{}, n int) {
	for k := 0; k < n; k++ {
		switch v := any.(type) {
		case *myint:
			v.inc()
		}
	}
}

func Benchmark_TypeAssertionInterface(b *testing.B) {
	i := new(myint)
	incTypeAssertionInterface(i, b.N)
}

func incTypeAssertionInterface(any interface{}, n int) {
	for k := 0; k < n; k++ {
		if newint, ok := any.(Inccer); ok {
			newint.inc()
		}
	}
}

func Benchmark_TypeAssertionPointer(b *testing.B) {
	i := new(myint)
	incAssertionPointer(i, b.N)
}

func incAssertionPointer(any interface{}, n int) {
	for k := 0; k < n; k++ {
		if newint, ok := any.(*myint); ok {
			newint.inc()
		}
	}
}
