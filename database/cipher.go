package database

import (
	"errors"
	"strconv"
	"time"

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
	if cipher.Id == "" {
		return errors.New("trying to update cipher without id (all ciphers with the same data)")
	}

	// construct a find user
	findCipher := Cipher{
		Id:    cipher.Id,
		Owner: cipher.Owner,
	}

	// find cipher and update it
	if err := DB.Where(&findCipher).Updates(cipher).Error; err != nil {
		return err
	}

	return nil
}

// Delete removes cipher.
func (cipher *Cipher) Delete() error {
	if err := DB.Delete(cipher).Error; err != nil {
		return err
	}

	return nil
}

// TakeOwnedCiphers returns all ciphers owned by user from the database.
func (user *User) TakeOwnedCiphers(lastSync string) (Ciphers, error) {
	var result Ciphers

	whereCondition := Cipher{
		Owner: user.Id,
	}

	if lastSync != "" && lastSync != "0" {
		unix, err := strconv.ParseInt(lastSync, 10, 64)
		if err != nil {
			return result, err
		}

		tx := DB.Unscoped().Scopes(UpdatedOrDeletedAfter(time.Unix(unix, 0))).Select("id", "updated_at", "deleted_at").Where(&whereCondition).Find(&result)
		return result, tx.Error
	} else {
		tx := DB.Select("id").Where(&whereCondition).Find(&result)
		return result, tx.Error
	}
}
