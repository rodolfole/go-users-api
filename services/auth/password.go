package auth

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
	return err == nil
}

func IsValidContrasena(contrasena string) bool {
	// Expresión regular para validar contraseña
	regex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$&])[A-Za-z\d@$&]{6,12}$`)
	return regex.MatchString(contrasena)
}

func IsValidPassword(password string) bool {
	// Requisitos de complejidad de contraseña
	hasUpper := false
	hasLower := false
	hasSpecial := false
	hasDigit := false

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case char == '@' || char == '$' || char == '&':
			hasSpecial = true
		case '0' <= char && char <= '9':
			hasDigit = true
		}
	}

	return len(password) >= 6 && len(password) <= 12 && hasUpper && hasLower && hasSpecial && hasDigit
}
