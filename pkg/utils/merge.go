package utils

import "github.com/luyasr/simple-blog/pkg/e"

func Merge(dest, src any) (map[string]any, error) {
	// 结构体转map
	m, err := Struct2Map(dest)
	if err != nil {
		return nil, err
	}

	// 获取非零字段
	fields, err := WithNotEmptyValue(src)
	if err != nil {
		return nil, err
	}

	// 去除目标与源map值相同字段
	for field := range fields {
		if m[field] == fields[field] {
			delete(fields, field)
		}
	}
	if len(fields) == 0 {
		return nil, e.NewUpdateFailed("当前记录未发生改变")
	}

	return fields, nil
}
