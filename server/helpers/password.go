package helpers

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/okoge-kaz/golang-todo-application/server/entities"
)

func CheckPasswordHash(user entities.User, password string) bool {
	return hex.EncodeToString(hash(password)) == hex.EncodeToString(user.Password)
}

func EncryptPassword(password string) []byte {
	return hash(password)
}

func hash(password string) []byte {
	const salt = "random_value_salt"

	h := sha256.New()
	h.Write([]byte(salt))
	h.Write([]byte(password))

	return h.Sum(nil)
}
