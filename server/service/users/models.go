package users

import (
	"errors"
	"mini-oa-server/common/database"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint   `gorm:"primary_key"`
	Username     string `gorm:"column:username"`
	Email        string `gorm:"column:email;uniqueIndex"`
	PasswordHash string `gorm:"column:password; not null"`
}

func AutoMigrate() {
	db := database.GetDB()
	db.AutoMigrate(&User{})
}

func FindOneUser(condition interface{}) (User, error) {
	db := database.GetDB()
	var model User
	err := db.Where(condition).First(&model).Error
	return model, err
}

func (u *User) Save() error {
	err := database.GetDB().Save(u).Error
	return err
}

func (u *User) SetPassword(password string) error {
	if len(password) <= 0 {
		return errors.New("password should not be empty")
	}

	bytesPwd := []byte(password)
	pwdHash, _ := bcrypt.GenerateFromPassword(bytesPwd, bcrypt.DefaultCost)
	u.PasswordHash = string(pwdHash)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}
