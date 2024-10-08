package orm

import (
	"reflect"
)

// The function GetType takes a reflect.Type as input and returns a string representing the
// corresponding SQL data type.
func GetType(fieldType reflect.Type) (sqlType string) {
	switch fieldType.Kind() {
	case reflect.Int64, reflect.Int:
		sqlType = "INTEGER"
	case reflect.String:
		sqlType = "TEXT"
	case reflect.Float64, reflect.Float32:
		sqlType = "REAL"
	case reflect.Struct:
		if fieldType.Name() == "Time" {
			sqlType = "DATETIME"
		}
	case reflect.Slice:
		if fieldType.Elem().Kind() == reflect.Uint8 {
			sqlType = "BLOB"
		}
	}
	return sqlType
}

// func CheckTypeSlice(fieldType reflect.Type) bool {
// 	for elem := range fieldType.Elem().Kind() {
// 		if elem != reflect.Uint8 {
// 			return false
// 		}
// 	}
// 	return true
// }
