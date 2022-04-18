package storage

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/vipindasvg/ticket-management/common"
)

const errnorec = "record not found"

type cursor struct {
	Db *gorm.DB
}

// cursor is used for interaction with the database
func GetCursor() *cursor {
	c := new(cursor)
	c.Db = common.Db
	return c
}
