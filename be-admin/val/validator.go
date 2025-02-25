package val

import (
	config "comics-admin/util"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func validateString(input string, minLength, maxLength int) error {

	n := len(input)

	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}

	return nil
}

var ValidateRole = func(fieldLevel validator.FieldLevel) bool {

	if role, ok := fieldLevel.Field().Interface().(string); ok {
		return config.IsSupportedRoles(role)
	}
	return false
}

func ValidateUsername(value string) error {
	if err := validateString(value, 3, 100); err != nil {
		return err
	}

	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lowercase letters, digits, or underscore")
	}

	return nil
}

func ValidateFullName(value string) error {
	if err := validateString(value, 3, 100); err != nil {
		return err
	}

	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letters or spaces")
	}
	return nil
}

func ValidatePassword(value string) error {
	return validateString(value, 6, 100)
}
func ValidateEmail(value string) error {

	if err := validateString(value, 3, 200); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("%s is not a valid email address", value)
	}
	return nil
}

func ValidateEmailId(value int) error {

	if value <= 0 {
		return fmt.Errorf("must be a positive integer")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return validateString(value, 32, 128)
}
