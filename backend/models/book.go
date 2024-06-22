package models

import (
    "gorm.io/gorm"
)

// Book 图书模型
type Book struct {
    gorm.Model
    BookInfo
}

type BookInfo struct {
    Title     string `json:"title" gorm:"type:varchar(255)" binding:"required"`
    Author    string `json:"author" gorm:"type:varchar(255)" binding:"required"`
    Category  string `json:"category" gorm:"type:varchar(100)" binding:"required"`
    Quantity  int    `json:"quantity" gorm:"type:int" binding:"required"`
}
