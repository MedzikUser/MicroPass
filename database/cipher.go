package database

import "fmt"

type Cipher struct {
	Model
	UserId         *string
	OrganizationId *string
	Type           int
	Data           string
	Collection     *string
}

type Ciphers = map[string]Cipher

var (
	CipherTypeLogin      = 1
	CipherTypeSecureNote = 2
	CipherTypeCard       = 3
)

// Insert cipher into the database.
func (cipher Cipher) Insert() (Cipher, error) {
	if cipher.UserId == nil && cipher.OrganizationId == nil {
		return cipher, fmt.Errorf("UserId or OrganizationId must be provided")
	}

	if cipher.UserId != nil && cipher.OrganizationId != nil {
		return cipher, fmt.Errorf("UserId and OrganizationId cannot be provided together, only one can be provided")
	}

	// add default entries to the cipher
	cipher.Model = defaultModel()

	tx := DB.Create(&cipher)

	return cipher, tx.Error
}

// Get cipher from the database
func (cipher Cipher) Get() (Cipher, error) {
	tx := DB.Model(&cipher).First(&cipher)

	return cipher, tx.Error
}

func (cipher Cipher) Update(id string) error {
	var findCipher Cipher
	findCipher.Id = id

	tx := DB.Model(&cipher).Where(&findCipher).Updates(&cipher)
	if tx.Error != nil {
		return fmt.Errorf("updating cipher in database error: %v", tx.Error)
	}

	return nil
}

// Get all user ciphers from the database.
func GetUserCiphers(userId string) (Ciphers, error) {
	var result Ciphers

	tx := DB.Model(&Cipher{UserId: &userId}).Take(&result)

	return result, tx.Error
}
