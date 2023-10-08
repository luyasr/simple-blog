package utils

import (
	"reflect"
)

func reflectObj(obj any) (reflect.Type, reflect.Value) {
	typeOf, valueOf := reflect.TypeOf(obj), reflect.ValueOf(obj)
	if typeOf.Kind() == reflect.Ptr {
		typeOf = reflect.TypeOf(obj).Elem()
	}
	if valueOf.Kind() == reflect.Ptr {
		valueOf = reflect.ValueOf(obj).Elem()
	}
	return typeOf, valueOf
}

func withNotEmptyValue(fields map[string]any, typeOf reflect.Type, valueOf reflect.Value) error {
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
			err := withNotEmptyValue(fields, tFiled.Type.Elem(), vFiled.Elem())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func WithNotEmptyValue(obj any) (map[string]any, error) {
	fields := make(map[string]any)
	typeOf, valueOf := reflectObj(obj)
	err := withNotEmptyValue(fields, typeOf, valueOf)
	if err != nil {
		return nil, err
	}
	return fields, nil
}

func struct2Map(result map[string]any, typeOf reflect.Type, valueOf reflect.Value) error {
	for i := 0; i < typeOf.NumField(); i++ {
		tFiled := typeOf.Field(i)
		vFiled := valueOf.Field(i)
		switch tFiled.Type.Kind() {
		case reflect.String:
			field, _ := tFiled.Tag.Lookup("json")
			result[field] = vFiled.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field, _ := tFiled.Tag.Lookup("json")
			result[field] = vFiled.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field, _ := tFiled.Tag.Lookup("json")
			result[field] = vFiled.Uint()
		case reflect.Ptr:
			err := struct2Map(result, tFiled.Type.Elem(), vFiled.Elem())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Struct2Map(obj any) (map[string]any, error) {
	result := make(map[string]any)
	typeOf, valueOf := reflectObj(obj)
	err := struct2Map(result, typeOf, valueOf)
	if err != nil {
		return nil, err
	}
	return result, nil
}
