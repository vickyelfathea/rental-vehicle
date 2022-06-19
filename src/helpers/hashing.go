package helpers

import ("golang.org/x/crypto/bcrypt"
		"fmt"
	)

func HashPassword(pass string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashPass), nil
}

func CheckPassword(hashPass, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))

	if err != nil {
		fmt.Println("password salah")
		return false
	}

	// fmt.Println("password salah")

	return true
}