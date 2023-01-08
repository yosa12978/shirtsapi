package helpers

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func NewMD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func NewUserToken() string {
	return strings.Replace(uuid.NewString(), "-", "", -1)
}

func NewSalt() string {
	buffer := make([]byte, 16)
	rand.Read(buffer)
	return fmt.Sprintf("%x", buffer)
}
