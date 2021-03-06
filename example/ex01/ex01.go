package main

import (
	"fmt"

	"github.com/teambition/validation"
)

// Person for test struct
type Person struct {
	Name     string   `valid:"required"`
	Email    string   `valid:"required;email"`
	Age      int      `valid:"-"`
	Sex      int      ``
	WebSites []string `valid:"url"`
}

func main() {
	web1 := "http://www.do1618.com"
	web2 := "www.baidu.com"

	person1 := &Person{
		Name:     "dave",
		Email:    "dwh0403@163.com",
		WebSites: []string{web1, web2},
	}

	validater := validation.NewValidation()
	res := validater.Validate(person1)

	if res {
		fmt.Println("Person1 validate succeed!")
	} else {
		fmt.Printf("Person1 validate failed. %s", validater.ErrMsg())
	}

	validater.Reset()
	person2 := &Person{
		Email:    "dwh0403@163.com",
		WebSites: []string{web1, web2},
	}

	res = validater.Validate(person2)

	if !res {
		fmt.Printf("Person2 validate failed. %s\n", validater.ErrMsg())
	} else {
		fmt.Println("Person2 validate succeed!")
	}
}
