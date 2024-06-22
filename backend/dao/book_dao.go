package dao

import (
    "books-manage-system/models"
    "books-manage-system/config"
    "gorm.io/gorm"
)

func CreateBook(book *models.Book) error {
    return config.DB.Create(book).Error
}

func UpdateBook(tx *gorm.DB, book *models.Book) error {
    if tx == nil {
        tx = config.DB
    }
    return tx.Save(book).Error
}

func DeleteBook(id uint) error {
    return config.DB.Delete(&models.Book{}, id).Error
}

func GetBookByID(tx *gorm.DB, id uint) (*models.Book, error) {
    var book models.Book
    if tx == nil {
        tx = config.DB
    }
    if err := tx.First(&book, id).Error; err != nil {
        return nil, err
    }
    return &book, nil
}

func GetBookList(page int, pageSize int) ([]*models.Book, int64, error) {
    var books []*models.Book
    var total int64

    db := config.DB.Model(&models.Book{})
    db.Count(&total)

    if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&books).Error; err != nil {
        return nil, 0, err
    }

    return books, total, nil
}

