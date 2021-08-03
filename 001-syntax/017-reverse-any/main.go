package main

import (
	"errors"
	"fmt"
	"reflect"
)

func reverseSlice(data interface{}) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		panic(errors.New("data must be a slice type"))
	}
	valueLen := value.Len()
	for i := 0; i <= int((valueLen-1)/2); i++ {
		reverseIndex := valueLen - 1 - i
		tmp := value.Index(reverseIndex).Interface()
		value.Index(reverseIndex).Set(value.Index(i))
		value.Index(i).Set(reflect.ValueOf(tmp))
		//fmt.Println(value) //reverse between
	}
	fmt.Println(value)
}

func main() {
	names := []string{"fua", "mua", "nata", "eva"}
	ages := []int{11, 22, 33, 44}

	reverseSlice(names)
	fmt.Println("---")
	reverseSlice(ages)
}
