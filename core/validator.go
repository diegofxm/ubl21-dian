package core

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidNIT    = errors.New("NIT inválido")
	ErrInvalidEmail  = errors.New("email inválido")
	ErrInvalidDate   = errors.New("fecha inválida")
	ErrInvalidAmount = errors.New("monto inválido")
)

// ValidateNIT valida un NIT colombiano
func ValidateNIT(nit string) error {
	if nit == "" {
		return ErrInvalidNIT
	}

	// Validación básica: solo números y guión
	matched, _ := regexp.MatchString(`^\d{9,10}(-\d)?$`, nit)
	if !matched {
		return ErrInvalidNIT
	}

	return nil
}

// ValidateEmail valida un email
func ValidateEmail(email string) error {
	if email == "" {
		return ErrInvalidEmail
	}

	matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	if !matched {
		return ErrInvalidEmail
	}

	return nil
}

// ValidateAmount valida que un monto sea positivo
func ValidateAmount(amount float64) error {
	if amount < 0 {
		return ErrInvalidAmount
	}
	return nil
}
