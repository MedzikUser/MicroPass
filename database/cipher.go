package database

import (
	"github.com/google/uuid"
)

type Cipher struct {
	Model
	Id    string `gorm:"size:40,primaryKey"`
	Owner string
	Data  string
}

type Ciphers = []Cipher

// Insert creates an new cipher in the database.
func (cipher *Cipher) Insert() error {
	// generate uuid
	cipher.Id = uuid.NewString()

	// create user in the database
	if err := DB.Create(cipher).Error; err != nil {
		return err
	}

	return nil
}

// Take finds cipher in the database.
func (cipher *Cipher) Take() error {
	// find cipher in the database
	if err := DB.Where(cipher).First(cipher).Error; err != nil {
		return err
	}

	return nil
}

// Update finds cipher and updates it.
func (cipher *Cipher) Update() error {
	// construct a find user
	findCipher := Cipher{
		Id: cipher.Id,
	}

	// find cipher and update it
	if err := DB.Where(findCipher).Updates(cipher).Error; err != nil {
		return err
	}

	return nil
}

func (cipher *Cipher) Delete() error {
	if err := DB.Delete(cipher).Error; err != nil {
		return err
	}

	return nil
}

// TakeOwnedCiphers returns all ciphers owned by user from the database.
func (user *User) TakeOwnedCiphers() (Ciphers, error) {
	var result Ciphers
	tx := DB.Where(&Cipher{Owner: user.Id}).Take(&result)

	return result, tx.Error
}
