package storage

import (
	"errors"

	"gitlab.com/vipindasvg/ticket-management/models"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates user record in the database
func (c *cursor) CreateUser(user *models.User) (*models.User, error) {
	if len(user.Password) < 10 {
		return nil, errors.New("non valid password provided")
	}
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.HashPassword = hpass
	result := c.Db.Create(user)
	if err := result.Error; err != nil {
		return nil, err
	}
	record := result.Value.(*models.User)
	record.Password = ""
	record.HashPassword = nil
	return record, nil
}

// Login user checking his email and password
func (c *cursor) Login(email, password string) (*models.User, error) {
	// Find the existing user
	user := new(models.User)
	if err := c.Db.Where("email=?", email).First(user).Error; err != nil {
		return nil, err
	}
	// validate password, compare hashed password taken from the database
	// and a password received
	err := bcrypt.CompareHashAndPassword(user.HashPassword, []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}
