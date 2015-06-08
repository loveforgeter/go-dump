package dump

import (
	"fmt"
	"io"
	"reflect"
)

var (
	trueBytes          = []byte("true")
	falseBytes         = []byte("false")
	varBytes           = []byte("var")
	space              = []byte(" ")
	equal              = []byte("=")
	comma              = []byte(",")
	colon              = []byte(":")
	quote              = []byte("\"")
	pointer            = []byte("*")
	address            = []byte("&")
	openBrace          = []byte("{")
	closeBrace         = []byte("}")
	openBracket        = []byte("(")
	closeBracket       = []byte(")")
	openSqureBracket   = []byte("[")
	closeSquareBracket = []byte("]")
	rightArrow         = []byte("->")
	leftArrow          = []byte("<-")
)

func Dump(v interface{}, w io.Writer) {
	value := reflect.ValueOf(v)
	// var v type = xxx
	w.Write(varBytes)
	w.Write(space)
	w.Write([]byte("v"))
	w.Write(space)
	w.Write([]byte(value.Type().String()))
	w.Write(space)
	w.Write(equal)
	w.Write(space)
	dumpAny(value, w)
}
func dumpAny(v reflect.Value, w io.Writer) {
	kind := v.Kind()
	switch kind {
	case reflect.Bool:
		dumpBool(v, w)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64:
		dumpNumber(v, w)
	case reflect.String:
		dumpString(v, w)
	case reflect.Slice:
		dumpSlice(v, w)
	case reflect.Map:
		dumpMap(v, w)
	case reflect.Chan:
		dumpChan(v, w)
	case reflect.Struct:
		dumpStruct(v, w)
	case reflect.Ptr:
		dumpPtr(v, w)
	case reflect.Interface:
		dumpInterface(v, w)
	default:
	}
}

func dumpBool(v reflect.Value, w io.Writer) {
	if v.Bool() {
		w.Write(trueBytes)
	} else {
		w.Write(falseBytes)
	}
}

func dumpNumber(v reflect.Value, w io.Writer) {
	w.Write([]byte(fmt.Sprint(v.Interface())))
}

func dumpString(v reflect.Value, w io.Writer) {
	w.Write([]byte(v.String()))
}

func dumpSlice(v reflect.Value, w io.Writer) {
}

func dumpMap(v reflect.Value, w io.Writer) {
	w.Write([]byte(v.Type().String()))
	w.Write(openBrace)
	for _, key := range v.MapKeys() {
		dumpAny(key, w)
		w.Write(colon)
		dumpAny(v.MapIndex(key), w)
		w.Write(comma)
	}
	w.Write(closeBrace)
}

func dumpChan(v reflect.Value, w io.Writer) {

}

func dumpStruct(v reflect.Value, w io.Writer) {
}

func dumpPtr(v reflect.Value, w io.Writer) {
}

func dumpInterface(v reflect.Value, w io.Writer) {
}

func dumpNil(v reflect.Value, w io.Writer) {
	w.Write([]byte("nil"))
}

func unpackValue(v reflect.Value) reflect.Value {
	if v.IsNil() {
		return v
	}
	return v.Elem()
}
