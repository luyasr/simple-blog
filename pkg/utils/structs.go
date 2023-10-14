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

func structToMap(result map[string]any, typeOf reflect.Type, valueOf reflect.Value) error {
	for i := 0; i < typeOf.NumField(); i++ {
		tFiled := typeOf.Field(i)
		vFiled := valueOf.Field(i)
		switch tFiled.Type.Kind() {
		case reflect.String:
			field, _ := tFiled.Tag.Lookup("json")
			result[field] = vFiled.Interface()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field, _ := tFiled.Tag.Lookup("json")
			result[field] = vFiled.Interface()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field, _ := tFiled.Tag.Lookup("json")
			result[field] = vFiled.Interface()
		case reflect.Map:
			field, _ := tFiled.Tag.Lookup("json")
			result[field] = vFiled.Interface()
		case reflect.Ptr:
			err := structToMap(result, tFiled.Type.Elem(), vFiled.Elem())
			if err != nil {
				return err
			}
		case reflect.Struct:
			continue
		}
	}
	return nil
}

func StructToMap(obj any) (map[string]any, error) {
	result := make(map[string]any)
	typeOf, valueOf := reflectObj(obj)
	err := structToMap(result, typeOf, valueOf)
	if err != nil {
		return nil, err
	}
	return result, nil
}
