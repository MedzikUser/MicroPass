package database

import (
	"strings"

	"github.com/MedzikUser/MicroPass/utils"
	"github.com/MedzikUser/libcrypto-go"
	"github.com/MedzikUser/libcrypto-go/hash"
	"github.com/google/uuid"
)

type User struct {
	Model
	Id               string `gorm:"size:40,primaryKey"`
	Email            string `gorm:"unique"`
	EmailConfirmed   bool
	Username         string
	Password         string
	PasswordSalt     []byte
	PasswordHint     *string
	TwoFactor        bool
	TwoFactorSecret  *string
	TwoFactorRecover *string
	EncryptionKey    string
}

// Insert creates an new user in the database.
func (user *User) Insert() error {
	// get password from the given user model
	givenPassword := user.Password
	if len(givenPassword) == 0 {
		return ErrPasswordEmpty
	}

	// generate salt
	salt, err := libcrypto.GenerateSalt(utils.Config.Crypto.Salt)
	if err != nil {
		return err
	}

	// add password salt to the user model
	user.PasswordSalt = salt

	// compute hash of the given password
	user.Password = hash.Pbkdf2Hash256(givenPassword, user.PasswordSalt, utils.Config.Crypto.Iterations)

	// get username from email
	user.Username = strings.Split(user.Email, "@")[0]

	// generate uuid
	user.Id = uuid.NewString()

	// create user in the database
	if err := DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// Take finds user in the database.
func (user *User) Take() error {
	givenPassword := user.Password

	// clear password from the given user model
	user.Password = ""

	// get user from database
	if err := DB.Where(user).First(user).Error; err != nil {
		return err
	}

	// match password if given
	if givenPassword != "" {
		// match the given password with the password saved in the database
		ok := hash.Pbkdf2Match256(user.Password, givenPassword, user.PasswordSalt, utils.Config.Crypto.Iterations)
		if !ok {
			return ErrPasswordMismatch
		}
	}

	return nil
}

// Update finds user and updates it.
func (user *User) Update() error {
	// construct a find user
	findUser := User{
		Id: user.Id,
	}

	// find user and update it
	if err := DB.Where(findUser).Updates(user).Error; err != nil {
		return err
	}

	return nil
}

// GenerateAccessToken generates access token for the user.
func (user *User) GenerateAccessToken() (string, error) {
	return utils.GenerateAccessToken(user.Id)
}

// GenerateRefreshToken generates refresh token for the user.
func (user *User) GenerateRefreshToken() (string, error) {
	return utils.GenerateRefreshToken(user.Id)
}
