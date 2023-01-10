package core

import "reflect"

func StructToMapStr(obj interface{}) map[string]string {
	data := make(map[string]string)

	v := reflect.ValueOf(obj).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tField := t.Field(i)
		vField := v.Field(i)
		tFieldTag := string(tField.Tag.Get("query"))
		if len(tFieldTag) > 0 {
			data[tFieldTag] = vField.String()
		} else {
			data[tField.Name] = vField.String()
		}
	}

	return data
}
