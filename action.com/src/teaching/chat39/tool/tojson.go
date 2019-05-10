package tool

import (
	"bytes"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

func ToJson(v interface{}) ([]byte, error) {
	var (
		b   bytes.Buffer
		err error
	)
	if err = writeAnyDataStruct(&b, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func writeAnyDataStruct(buf *bytes.Buffer, value reflect.Value) error {
	switch value.Kind() {
	case reflect.String:
		buf.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		buf.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Slice:
		return writeSlice(buf, value)
	case reflect.Struct:
		return writeStrut(buf, value)
	default:
		return errors.New("unknown kind: " + value.Kind().String())
	}
	return nil
}

func writeSlice(buf *bytes.Buffer, value reflect.Value) error {
	buf.WriteString("[")
	for i := 0; i < value.Len(); i++ {
		v := value.Index(i)
		_ = writeAnyDataStruct(buf, v)
		if i < value.Len()-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString("]")
	return nil
}

func writeStrut(buf *bytes.Buffer, value reflect.Value) error {
	vt := value.Type()
	buf.WriteString("{")
	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldType := vt.Field(i)
		buf.WriteString(`"`)
		buf.WriteString(fieldType.Name)
		buf.WriteString(`":`)
		_ = writeAnyDataStruct(buf, fieldValue)
		if i < value.NumField()-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString("}")
	return nil
}
