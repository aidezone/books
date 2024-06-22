package models

import (
    "gorm.io/gorm"
    "books-manage-system/utils"
)

type User struct {
    gorm.Model
    Username string     `json:"username" gorm:"type:varchar(100);uniqueIndex"`
    Password string     `json:"password" gorm:"type:varchar(100)"`
    Role     utils.Role `json:"role" gorm:"type:varchar(20)"`
    Locked   bool       `json:"locked" gorm:"type:smallint"`
}
