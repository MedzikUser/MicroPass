package database

import (
	"fmt"
	"strings"
)

type User struct {
	Model
	Name               string
	Email              string `gorm:"unique"`
	EmailVerified      bool
	MasterPassword     string
	MasterPasswordSalt []byte
	MasterPasswordHint *string
	TwoFactorSecret    *string
	TwoFactorRecover   *string
}

func NewUser(email string, masterPassword string, masterPasswordHint string) (User, error) {
	masterPasswordSalt, err := generateSalt()
	if err != nil {
		return User{}, err
	}

	hashedMasterPassword := hashPassword(masterPassword, masterPasswordSalt)

	name := strings.Split(email, "@")[0]

	user := User{
		Model:              defaultModel(),
		Name:               name,
		Email:              email,
		MasterPassword:     hashedMasterPassword,
		MasterPasswordSalt: masterPasswordSalt,
		MasterPasswordHint: &masterPasswordHint,
	}

	tx := DB.Create(&user)

	return user, tx.Error
}

func TakeUser(email string, masterPassword string) (*User, error) {
	var user User

	tx := DB.Where(map[string]interface{}{"email": email}).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	match := PasswordMatch(user.MasterPassword, masterPassword, user.MasterPasswordSalt)
	if !match {
		return nil, fmt.Errorf("password mismatch")
	}

	return &user, nil
}

func TakeUserID(id string) (*User, error) {
	var user User

	tx := DB.Where(map[string]interface{}{"uuid": id}).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}
