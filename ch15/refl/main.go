package main

import (
	"reflect"
)

func main() {
	type Person struct {
		name string
		age  int
	}
	var per Person
	per.name = "张"
	per.age = 11

}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			data[t.Field(i).Name] = f.String()
		case reflect.Int:
			data[t.Field(i).Name] = f.Int()
		}
	}
	return data
}
