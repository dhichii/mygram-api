package helper

import "golang.org/x/crypto/bcrypt"

func CreateHash(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func ValidateHash(password, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashed))
	return err == nil
}
