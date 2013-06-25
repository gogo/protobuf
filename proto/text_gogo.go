package proto

import (
	"fmt"
	"reflect"
)

func writeEnum(w *textWriter, v reflect.Value, props *Properties) error {
	m, ok := enumStringMaps[props.Enum]
	if !ok {
		if err := writeAny(w, v, props); err != nil {
			return err
		}
	}
	key := int32(0)
	if v.Kind() == reflect.Ptr {
		key = int32(v.Elem().Int())
	} else {
		key = int32(v.Int())
	}
	s, ok := m[key]
	if !ok {
		if err := writeAny(w, v, props); err != nil {
			return err
		}
	}
	_, err := fmt.Fprint(w, s)
	return err
}
