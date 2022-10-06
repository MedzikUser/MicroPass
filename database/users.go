package database

import (
	"fmt"
	"strings"

	"github.com/bytepass/server/crypto"
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

func NewUser(email string, masterPassword string, masterPasswordHint string) (*User, error) {
	// generate salt
	masterPasswordSalt, err := crypto.GenerateSalt()
	if err != nil {
		return nil, err
	}

	// hash master password with salt
	hashedMasterPassword := crypto.HashPassword(masterPassword, masterPasswordSalt)

	// get username from email
	name := strings.Split(email, "@")[0]

	user := User{
		Model:              defaultModel(),
		Name:               name,
		Email:              email,
		MasterPassword:     hashedMasterPassword,
		MasterPasswordSalt: masterPasswordSalt,
		MasterPasswordHint: &masterPasswordHint,
	}

	// create user in database
	tx := DB.Create(&user)

	return &user, tx.Error
}

func TakeUser(email string, masterPassword string) (*User, error) {
	var user User

	// find user in database
	tx := DB.Where(map[string]interface{}{"email": email}).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// match the master password with the master password from database
	match := crypto.PasswordMatch(user.MasterPassword, masterPassword, user.MasterPasswordSalt)
	if !match {
		return nil, fmt.Errorf("password mismatch")
	}

	return &user, nil
}

func TakeUserID(id string) (*User, error) {
	var user User

	// find user in database using uuid
	tx := DB.Where(map[string]interface{}{"uuid": id}).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}
