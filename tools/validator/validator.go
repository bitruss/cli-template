package tools

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func ValidateEmail(email string) bool {
	validate := validator.New()
	err := validate.Var(email, "email")
	if err != nil {
		return false
	}
	return true
}

func ValidatePassword(password string) bool {
	//must contain number and letter, special character is optional,length 6-20
	if len(password) < 6 || len(password) > 20 {
		return false
	}
	var hasNumber, hasLetter bool
	for _, c := range password {
		if hasNumber && hasLetter {
			return true
		}
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsLetter(c):
			hasLetter = true
		}
	}
	return hasNumber && hasLetter
}
