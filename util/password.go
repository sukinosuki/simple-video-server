package util

import "golang.org/x/crypto/bcrypt"

type password struct {
}

var Password = &password{}

func (p *password) Hashed(originPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(originPassword), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (p *password) Compare(hashedPassword, originPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(originPassword))

	return err
}
