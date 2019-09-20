package byte_string_conventer

import (
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"time"
)
import "unsafe"

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var strTestLen100 = RandStringRunes(100)
var strTestLen10000 = RandStringRunes(10000)

func Test_ByteString(t *testing.T) {
	var x = []byte("should equal")
	var y = *(*string)(unsafe.Pointer(&x))
	var z = string(x)

	if y != z {
		t.Fail()
	}
}

func normal(b []byte, n int) {
	for i := 0; i < n; i++ {
		_ = string(b)
	}
}

func ByteStringStringsBuilder(b []byte) (s string) {
	s = *(*string)(unsafe.Pointer(&b))
	return
}

//strings.Builder(https://golang.org/src/strings/builder.go#L45)
func stringsBuilder(b []byte, n int) {
	for i := 0; i < n; i++ {
		_ = ByteStringStringsBuilder(b)
	}
}

func ByteStringKeepAlive(b []byte) (s string) {
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	str.Data = slice.Data
	str.Len = slice.Len
	runtime.KeepAlive(&b)
	return
}

func byteStringKeepAlive(b []byte, n int) {
	for i := 0; i < n; i++ {
		_ = ByteStringKeepAlive(b)
	}
}

func ByteStringKeepAliveOff(bytes []byte) (s string) {
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	str.Data = uintptr(unsafe.Pointer(&bytes[0]))
	str.Len = len(bytes)
	return
}

func byteStringKeepAliveOff(b []byte, n int) {
	for i := 0; i < n; i++ {
		_ = ByteStringKeepAliveOff(b)
	}
}

func Benchmark_Normal_Len_100(b *testing.B) {
	normal([]byte(strTestLen100), b.N)
}

func Benchmark_ByteStringStringsBuilder_Len_100(b *testing.B) {
	stringsBuilder([]byte(strTestLen100), b.N)
}

func Benchmark_ByteStringKeepAlive_Len_100(b *testing.B) {
	byteStringKeepAlive([]byte(strTestLen100), b.N)
}

func Benchmark_ByteStringKeepAliveOff_Len_100(b *testing.B) {
	byteStringKeepAliveOff([]byte(strTestLen100), b.N)
}

func Benchmark_Normal_Len_10000(b *testing.B) {
	normal([]byte(strTestLen10000), b.N)
}

func Benchmark_ByteStringStringsBuilder_Len_10000(b *testing.B) {
	stringsBuilder([]byte(strTestLen10000), b.N)
}

func Benchmark_ByteStringKeepAlive_Len_10000(b *testing.B) {
	byteStringKeepAlive([]byte(strTestLen10000), b.N)
}

func Benchmark_ByteStringKeepAliveOff_Len_10000(b *testing.B) {
	byteStringKeepAliveOff([]byte(strTestLen10000), b.N)
}
