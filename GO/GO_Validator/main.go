package main

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type RegisterReq struct {
	// gt = greater than
	Username    string `validate:"gt=3"`
	PasswordNew string `validate:"gt=3,password"`
	// eqfield = PasswordRepeat equal PasswordNew
	PasswordRepeat string `validate:"eqfield=PasswordNew"`
	// email format
	Email string `validate:"email"`
}

var validate = validator.New()

func validateFunc(req RegisterReq) error {
	err := validate.Struct(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("all clear")
	return err
}

func validatePassword(fl validator.FieldLevel) bool {
	if !strings.ContainsAny("0123456789", fl.Field().String()) || !strings.ContainsAny("~!@#$%^&*-_|;:,.<>[]{}()", fl.Field().String()) || len(fl.Field().String()) < 8 {
		return false
	}
	return true
}

func main() {
	validate.RegisterValidation("password", validatePassword)
	var req RegisterReq
	req.Username = "1"
	fmt.Println(req)
	// PasswordNew = PasswordRepeat = ""
	validateFunc(req)

	req.Username = "1234"
	fmt.Println(req)
	validateFunc(req)
	req.PasswordNew = "11234"
	req.Email = "example@example.com"
	fmt.Println(req)
	validateFunc(req)
	req.PasswordNew = "~11234#11"
	req.PasswordRepeat = "~11234#11"
	fmt.Println(req)
	validateFunc(req)

	// if strings.ContainsAny(numbers, req.Username) {
	// 	fmt.Println("A")
	// }
	// if strings.ContainsAny(special, req.Username) {
	// 	fmt.Println("B")
	// }

}
