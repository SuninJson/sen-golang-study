package main

import "reflect"

func IsNil(i interface{}) bool {
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		return reflect.ValueOf(i).IsNil()
	}
	return i == nil
}
