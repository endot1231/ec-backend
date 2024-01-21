package secrets

import "golang.org/x/crypto/bcrypt"

func Encrypt(value string) (string, error) {
	encryptS, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encryptS), nil
}

func CompareValue(encryptedVlalue string, value string) error {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedVlalue), []byte(value))
	return err
}
