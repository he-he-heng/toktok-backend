package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword 함수는 bcrypt를 사용하여 입력 비밀번호를 해싱합니다.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// ComparePassword 함수는 입력 비밀번호와 해시된 비밀번호를 비교합니다.
func ComparePassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
