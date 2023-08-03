package utils

import "reflect"

// 实现类似 Gin 的 Bind 效果，将 Request 中的 Query 从结构体转为 map
// 以便在生成发起请求时，使用 req.URL.Query().Add() 注意为请求中的 Request 添加 Query
// 这个功能好像只有在自己暴露 API，并且传入的参数需要当做发起其他请求的 Query 时才有用
func StructToMapStr(obj interface{}) map[string]string {
	data := make(map[string]string)

	v := reflect.ValueOf(obj).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tField := t.Field(i)
		vField := v.Field(i)
		// 注意！！！注意！！！注意！！！
		// 传入的结构体中，要带有 form Tag 的才可以被解析为 map
		// 使用 form 这个 Tag 的原因是 Gin 的转换 map 逻辑中，也是使用的 form 作为 Tag
		tFieldTag := string(tField.Tag.Get("form"))
		if len(tFieldTag) > 0 {
			data[tFieldTag] = vField.String()
		} else {
			data[tField.Name] = vField.String()
		}
	}

	return data
}
