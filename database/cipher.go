package database

import (
	"time"

	"github.com/google/uuid"
)

type Cipher struct {
	Model
	Id          string `gorm:"size:40,primaryKey"`
	Owner       string
	Favorite    bool
	Directory   *string
	Data        string
	Attachments []string `gorm:"type:text[]"`
}

// TODO: clear cipher data from database
// func (cipher *Cipher) BeforeDelete(tx *gorm.DB) (err error) {
// 	// delete data from the cipher
// 	cipher.Data = ""

// 	return nil
// }

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
		return ErrUpdateCipherEmptyID
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
func (user *User) TakeOwnedCiphers(lastSync *int64) (Ciphers, error) {
	// construct a WHERE clause
	whereCondition := Cipher{
		Owner: user.Id,
	}

	var result Ciphers

	if lastSync != nil && *lastSync != 0 {
		// get all user ciphers that were updated after last sync time
		tx := DB.Unscoped().Scopes(UpdatedOrDeletedAfter(time.Unix(*lastSync, 0))).Select("id", "updated_at", "deleted_at").Where(&whereCondition).Find(&result)
		return result, tx.Error
	} else {
		// get all ciphers owned by the user
		tx := DB.Select("id").Where(&whereCondition).Find(&result)
		return result, tx.Error
	}
}
