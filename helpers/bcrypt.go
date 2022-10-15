package helpers

import "golang.org/x/crypto/bcrypt"

func HasPass(pass string) string {
	salt := 8

	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), salt)
	return string(hash)

}

func ComparePass(h, p []byte) bool {
	has, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(has, pass)
	return err == nil
}
