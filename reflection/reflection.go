package reflection

import "reflect"

// > golang challenge: write a function `walk(x interface{}, fn func(string))` which takes a struct `x` and calls `fn` for all strings fields found inside. difficulty level: recursively.

func Walk(x any, fn func(string)) {
	val := getValue(x)

    if val.Kind() == reflect.Slice {
        for i := 0; i < val.Len(); i++ {
            Walk(val.Index(i).Interface(), fn)
        }
        return
    }

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			Walk(field.Interface(), fn)
		}
	}
}

func getValue(x any) reflect.Value {
    val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

    return val
}
