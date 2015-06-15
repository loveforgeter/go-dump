package dump

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
)

var (
	trueBytes          = []byte("true")
	falseBytes         = []byte("false")
	varBytes           = []byte("var")
	space              = []byte(" ")
	equal              = []byte("=")
	comma              = []byte(",")
	colon              = []byte(":")
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

func Dump(v interface{}, w io.Writer, name ...string) {
	value := reflect.ValueOf(v)
	// var v type = name
	varName := "v"
	if len(name) != 0 {
		varName = name[0]
	}
	w.Write(varBytes)
	w.Write(space)
	w.Write([]byte(varName))
	w.Write(space)
	w.Write([]byte(value.Type().String()))
	w.Write(space)
	w.Write(equal)
	w.Write(space)
	dumpAny(value, w)
}

func dumpAny(v reflect.Value, w io.Writer) {
	if !v.CanInterface() {
		return
	}

	kind := v.Kind()
	switch kind {
	case reflect.Bool:
		dumpBool(v, w)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64:
		dumpNumber(v, w)
	case reflect.String:
		dumpString(v, w)
	case reflect.Slice, reflect.Array:
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
		panic(fmt.Sprint("unknown type %v", kind))
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
	w.Write([]byte(strconv.Quote(v.String())))
}

func dumpSlice(v reflect.Value, w io.Writer) {
	w.Write(openSqureBracket)
	for i := 0; i < v.Len(); i++ {
		dumpAny(v.Index(i), w)
		if i != v.Len()-1 {
			w.Write(comma)
		}
	}
	w.Write(closeSquareBracket)
}

func dumpMap(v reflect.Value, w io.Writer) {
	w.Write([]byte(v.Type().String()))
	w.Write(openBrace)
	for i, key := range v.MapKeys() {
		dumpAny(key, w)
		w.Write(colon)
		dumpAny(v.MapIndex(key), w)
		if i != len(v.MapKeys())-1 {
			w.Write(comma)
		}
	}
	w.Write(closeBrace)
}

func dumpChan(v reflect.Value, w io.Writer) {
}

func dumpStruct(v reflect.Value, w io.Writer) {
	t := v.Type()
	w.Write([]byte(t.Name()))
	w.Write(openBrace)
	numFields := v.NumField()
	for i := 0; i < numFields; i++ {
		if !v.Field(i).CanInterface() {
			continue
		}
		w.Write([]byte(t.Field(i).Name))
		w.Write(colon)
		dumpAny(v.Field(i), w)
		w.Write(comma)
	}
	w.Write(closeBrace)
}

func dumpPtr(v reflect.Value, w io.Writer) {
	if v.IsNil() {
		return
	}

	vt := v.Type()
	for vt.Kind() == reflect.Ptr {
		v = v.Elem()
		vt = v.Type()
	}
	dumpAny(v, w)
}

func dumpInterface(v reflect.Value, w io.Writer) {
}

func dumpNil(v reflect.Value, w io.Writer) {
	w.Write([]byte("nil"))
}
