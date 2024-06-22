package services

import (
    "books-manage-system/models"
    "books-manage-system/dao"
    "books-manage-system/utils"
    "golang.org/x/crypto/bcrypt"
    "errors"
)

func Register(username, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := &models.User{
        Username: username,
        Password: string(hashedPassword),
        Role:     utils.User,
    }

    return dao.CreateUser(user)
}

func Login(username, password string) (*models.User, error) {
    user, err := dao.GetUserByUsername(username)
    if err != nil {
        return nil, err
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, errors.New("invalid username or password")
    }

    return user, nil
}
