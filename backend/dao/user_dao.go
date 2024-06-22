package dao

import (
    "books-manage-system/models"
    "books-manage-system/config"
)

func CreateUser(user *models.User) error {
    return config.DB.Create(user).Error
}

func GetUserByUsername(username string) (*models.User, error) {
    var user models.User
    err := config.DB.Where("username = ?", username).First(&user).Error
    return &user, err
}

func GetUserByID(id uint) (*models.User, error) {
    var user models.User
    err := config.DB.First(&user, id).Error
    return &user, err
}

func LockOrUnlockUser(id uint, locker bool) error {
    return config.DB.Model(&models.User{}).Where("id = ?", id).Update("locked", locker).Error
}

func GetUsers(page int, pageSize int) ([]*models.User, int64, error) {
    var users []*models.User
    var total int64

    db := config.DB.Model(&models.User{})
    db.Count(&total)

    if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
        return nil, 0, err
    }

    return users, total, nil
}