package orm

import (
	"database/sql"
	"reflect"
)

var (
	__BUILDER__ = NewSQLBuilder()
	__QUERY__   string
	__PARAMS__  []interface{}
)

// The `Insert` function is a method of the `ORM` struct. It takes in one or more tables as arguments,
// which are of type `interface{}`.
func (o *ORM) Insert(tables ...interface{}) ([]sql.Result, error) {
	var results []sql.Result
	for _, t := range tables {

		_, __TABLE__ := InitTable(t)

		if reflect.TypeOf(t).Kind() == reflect.Struct {
			var values []interface{}
			v := reflect.ValueOf(t)

			for i := 0; i < v.NumField(); i++ {
				switch v.Field(i).Kind() {
				case reflect.Int64, reflect.Int:
					values = append(values, v.Field(i).Int())
				case reflect.String:
					values = append(values, v.Field(i).String())
                case reflect.Slice:
					values = append(values, v.Field(i).Bytes())
				case reflect.Float64, reflect.Float32:
					values = append(values, v.Field(i).Float())
				case reflect.Struct:
					if v.Field(i).Type().Name() == "Model" {
						__TABLE__.AllFields = __TABLE__.AllFields[2:]
					}
				}
			}


			if len(values) > 0 {
				__BUILDER__.Clear()
				__QUERY__, __PARAMS__ = __BUILDER__.Insert(__TABLE__, values).Build()
				rlt, err := o.Db.Exec(__QUERY__, __PARAMS__...)
				if err != nil {
					return nil, err
				}
				results = append(results, rlt)
			}

		}
	}
	return results, nil
}
