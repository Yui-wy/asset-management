package util

import (
	"crypto/md5"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func CreateMD5Random(username string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%.6f", username, rand.Float32()+1))))
}
