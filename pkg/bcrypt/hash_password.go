package pkgbcrypt

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	hashedByte, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(hashedByte), nil
}

func CheckPasswordHash(password, hasedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))

	return err == nil
}