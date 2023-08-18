package util

import (
	"reflect"
	"strconv"
	"strings"
)

type MyStruct struct {
	Field1 string
	Field2 int
	Field3 float64
	Field4 bool
	Field5 string
}

func StructToMap(obj interface{}, excludeEmptyFields bool) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(obj)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// 如果指定排除空字段，并且当前字段值为空，则跳过该字段
		if excludeEmptyFields && IsZeroValue(value) {
			continue
		}

		result[field.Name] = value.Interface()
	}

	return result
}

func IsZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	default:
		return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
	}
}

// 下划线转大写驼峰
func UnderscoreToCamelCase(underscore string) string {
	words := strings.Split(underscore, "_")
	for i := 0; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}




func Str2defaultNum(s string, defauleNum int) (n int) {
	var err error
	n, err = strconv.Atoi(s)
	if err != nil {
		n = defauleNum
	}
	return n
}
