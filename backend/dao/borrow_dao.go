package dao

import (
    "books-manage-system/models"
    "books-manage-system/config"
    "gorm.io/gorm"
)

func CreateBorrow(tx *gorm.DB, borrow *models.Borrow) error {
    return tx.Create(borrow).Error
}

func GetBorrowByID(tx *gorm.DB, id uint) (models.Borrow, error) {
    var borrow models.Borrow
    if err := tx.First(&borrow, id).Error; err != nil {
        return borrow, err
    }
    return borrow, nil
}

func UpdateBorrow(tx *gorm.DB, borrow models.Borrow) error {
    return tx.Save(&borrow).Error
}

func GetBorrowedBooks(userID uint, page int, pageSize int) ([]*models.Borrow, int64, error) {
    var borrows []*models.Borrow
    var total int64

    db := config.DB.Model(&models.Borrow{}).Where("user_id = ?", userID)
    db.Count(&total)

    if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&borrows).Error; err != nil {
        return nil, 0, err
    }

    return borrows, total, nil
}