package storage

import (
	"errors"
	"fmt"

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
	//fmt.Println("user in createuser function: ", user)
	fmt.Println("user.is_admin :", user.IsAdmin.Bool)
	result := c.Db.Create(user)
	if err := result.Error; err != nil {
		return nil, err
	}
	record := result.Value.(*models.User)
	record.Password = ""
	record.HashPassword = nil
	return record, nil
}

//create ticket
func (c *cursor) CreateTicket(user *models.Ticket) (*models.Ticket, error) {

	fmt.Println("Inside create ticket in storage:", user)
	result := c.Db.Create(user)
	fmt.Println("result : ", result)
	if err := result.Error; err != nil {
		return nil, err
	}

	record := result.Value.(*models.Ticket)
	fmt.Println("record:", record)
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

//List user
func (c *cursor) UserList(psql string) ([]models.User, error) {
	var data []models.User

	if err := c.Db.Raw(psql).Scan(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

//Get all tickets
func (c *cursor) TicketList(psql string) ([]models.Ticket, error) {
	var data []models.Ticket

	if err := c.Db.Raw(psql).Scan(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

//List all tickets created by user
func (c *cursor) ListUserTickets(psql string) ([]models.Ticket, error) {
	var data []models.Ticket

	if err := c.Db.Raw(psql).Scan(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

//List user ticket detail
func (c *cursor) UserTicketDetail(psql string) ([]models.Ticket, error) {
	var data []models.Ticket

	if err := c.Db.Raw(psql).Scan(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

//update tickets
func (c *cursor) UpdateUserTicket(t *models.Ticket, tid, uid int) error {
	var data []models.Ticket
	result := c.Db.Model(&data).Where("ticket_id = ? AND u_id = ?", tid, uid).Updates(
		models.Ticket{
			Desk:        t.Desk,
			Ticket_Type: t.Ticket_Type,
			Subject:     t.Subject,
		})
	if result.RowsAffected != 1 {
		return fmt.Errorf("Access deinied.. You don't have permission to update tickets created by other user..")
	}

	return nil
}
