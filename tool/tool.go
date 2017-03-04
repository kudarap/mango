package tool

import "reflect"

// InArray check if needle exist on haystack
func InArray(val interface{}, array interface{}) (index int, exists bool) {
	index = -1
	exists = false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
