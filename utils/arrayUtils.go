package utils

import (
	"reflect"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/8 12:13 AM
 */
func Contarins(val interface{}, array []interface{}) bool {

	switch reflect.TypeOf(array).Kind() {

	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {

			if val == s.Index(i).Interface() {
				return true

			}
		}
		return false

	}
	return false

}
