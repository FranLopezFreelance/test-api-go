package db

import "golang.org/x/crypto/bcrypt"

// PasswordEncrypt encripta las passqord
func PasswordEncrypt(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}