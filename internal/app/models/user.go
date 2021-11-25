package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id                int
	Email             string
	Username          string
	EncryptedPassword string
}

func (u *User) HashPassword() error {
	enc, err := encryptString(u.EncryptedPassword)
	if err != nil {
		return err
	}

	u.EncryptedPassword = enc
	return nil
}

func encryptString(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
