package models

import (
    "gorm.io/gorm"
    "time"
)

// Borrow 借阅记录模型
type Borrow struct {
    gorm.Model
    UserID   uint      `json:"user_id" gorm:"not null"`
    BookID   uint      `json:"book_id" gorm:"not null"`
    ReturnedAt *time.Time `json:"returned_at" gorm:"default null"`
}
