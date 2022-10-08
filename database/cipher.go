package database

type Cipher struct {
	Model
	UserId         *string
	OrganizationId *string
	Type           int
	Data           string
	Collection     *[]string `gorm:"type:text[]"`
}

func (cipher Cipher) Create() (Cipher, error) {
	// add default entries to the cipher
	cipher.Model = defaultModel()

	// insert cipher into database
	tx := DB.Create(&cipher)

	return cipher, tx.Error
}

func TakeUserCiphers(userId string) {
	var result map[string]Cipher

	DB.Model(&Cipher{UserId: &userId}).Take(&result)
}
