package struct2mapstrstr

import (
	"fmt"
	"reflect"
)

func ConvertMapStringString(s interface{}) map[string]string {
	r := map[string]string{}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		k := t.Field(i).Tag.Get("mss")
		if k != "" {
			f := v.Field(i)
			switch f.Kind() {
			case reflect.Ptr:
				if !f.IsNil() {
					r[k] = fmt.Sprintf("%v", f.Elem().Interface())
				}
			default:
				r[k] = fmt.Sprintf("%v", f.Interface())

			}
		}
	}
	return r
}
