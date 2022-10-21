package database

import (
	"fmt"
	"strings"

	"github.com/bytepass/libcrypto-go"
	"github.com/bytepass/server/config"
)

// User table in database.
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

// Insert user into the database.
func (user User) Insert() (User, error) {
	// generate salt
	salt, err := libcrypto.GenerateSalt(config.Config.Crypto.Salt)
	if err != nil {
		return user, fmt.Errorf("generate salt error: %v", err)
	}
	// set user salt to the generated salt
	user.MasterPasswordSalt = salt

	// compute hash from master password with salt
	user.MasterPassword = libcrypto.Pbkdf2Hash256(user.MasterPassword, user.MasterPasswordSalt, config.Config.Crypto.Iterations)

	// get username from email
	user.Name = strings.Split(user.Email, "@")[0]

	// add default entries to the user
	user.Model = defaultModel()

	// create user in the database
	tx := DB.Create(&user)
	if tx.Error != nil {
		return user, fmt.Errorf("inseting user to database error: %v", tx.Error)
	}

	return user, nil
}

// Get user from the database.
func (user User) Get() (User, error) {
	providedMasterPassword := user.MasterPassword
	user.MasterPassword = ""

	// get user from database
	tx := DB.Model(&user).Where(&user).First(&user)
	if tx.Error != nil {
		return user, tx.Error
	}

	// validate password if provided
	if providedMasterPassword != "" {
		// match the master password with the master password saved in the database
		ok := libcrypto.Pbkdf2Match256(user.MasterPassword, providedMasterPassword, user.MasterPasswordSalt, config.Config.Crypto.Iterations)
		if !ok {
			return user, fmt.Errorf("password mismatch")
		}
	}

	return user, nil
}

// Update user data in the database.
func (user User) Update(id string) error {
	var findUser User
	findUser.Id = id

	tx := DB.Model(&user).Where(&findUser).Updates(&user)
	if tx.Error != nil {
		return fmt.Errorf("updating user in database error: %v", tx.Error)
	}

	return nil
}
