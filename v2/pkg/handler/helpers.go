package handler

import (
	"encoding/base64"
	"strings"
	"unicode"
)

func MaybeDecode(value string) string {
	if strings.HasPrefix(value, "base64: ") {
		udec, err := base64.StdEncoding.DecodeString(value[8:])
		if err != nil {
			value = "* invalid value *"
		} else {
			value = string(udec)
		}
	}
	return value
}

func validatePassword(password string) bool {
	// Check minimum length requirement
	if len(password) < 14 {
		return false
	}

	// Check for at least one uppercase letter
	hasUppercase := false
	// Check for at least one lowercase letter
	hasLowercase := false
	// Check for at least one digit
	hasDigit := false
	// Check for at least one special character
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	// Check if all required character types are present
	return hasUppercase && hasLowercase && hasDigit && hasSpecial
}
