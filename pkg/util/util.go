package util

import (
	"reflect"
)

func subUpdateNonZeroFields(fields map[string]any, typeOf reflect.Type, valueOf reflect.Value) error {
	for i := 0; i < typeOf.NumField(); i++ {
		tFiled := typeOf.Field(i)
		vFiled := valueOf.Field(i)
		switch tFiled.Type.Kind() {
		case reflect.String:
			field, _ := tFiled.Tag.Lookup("json")
			if vFiled.String() != "" {
				fields[field] = vFiled.String()
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field, _ := tFiled.Tag.Lookup("json")
			if vFiled.Int() != 0 {
				fields[field] = vFiled.Int()
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field, _ := tFiled.Tag.Lookup("json")
			if vFiled.Uint() != 0 {
				fields[field] = vFiled.Uint()
			}
		case reflect.Ptr:
			err := subUpdateNonZeroFields(fields, tFiled.Type.Elem(), vFiled.Elem())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func UpdateNonZeroFields(obj any) (map[string]any, error) {
	fields := make(map[string]any)
	typeOf, valueOf := reflect.TypeOf(obj).Elem(), reflect.ValueOf(obj).Elem()
	err := subUpdateNonZeroFields(fields, typeOf, valueOf)
	if err != nil {
		return nil, err
	}
	return fields, nil
}
