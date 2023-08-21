package validator

import (
	"fmt"
	"reflect"
)

type Data struct {
	ID    int    `required:"true" max:"2147483647"`
	Name  string `maxlen:"255"`
	Age   int    `min:"18" max:"2000"`
	Phone string `minlen:"3" maxlen:"31"`
	Email string `email:"true" required:"true"`
}

func (d Data) Validate() {

	refType := reflect.TypeOf(d)
	refVal := reflect.ValueOf(d)
	for i := 0; i < refType.NumField(); i++ {
		x := refType.Field(i)
		interfaceVal := refVal.Field(i).Interface()
		if x.Tag.Get("required") == "true" && interfaceVal == "" {
			fmt.Println("Field", x.Name, "is empty.")
		} else {
			fmt.Println("ok")
		}
	}
}
