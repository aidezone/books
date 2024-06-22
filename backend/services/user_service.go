package services

import (
    "books-manage-system/models"
    "books-manage-system/dao"
    "books-manage-system/utils"
    "strconv"
    "os"
    "log"
    "golang.org/x/crypto/bcrypt"
)

// 锁定用户
func LockUser(id string) error {
    userID, err := strconv.Atoi(id)
    if err != nil {
        return err
    }
    return dao.LockOrUnlockUser(uint(userID), true)
}

// 解锁用户
func UnlockUser(id string) error {
    userID, err := strconv.Atoi(id)
    if err != nil {
        return err
    }
    return dao.LockOrUnlockUser(uint(userID), false)
}

// 获取所有用户
func GetUsers(page int, pageSize int) ([]*models.User, int64, error) {
    return dao.GetUsers(page, pageSize)
}

// 初始化管理员账号
func InitializeAdmin() {
    adminUsername := os.Getenv("ADMIN_USERNAME")
    adminPassword := os.Getenv("ADMIN_PASSWORD")

    if adminUsername == "" || adminPassword == "" {
        log.Println("Admin credentials are not set in environment variables")
        return
    }

    existingAdmin, err := dao.GetUserByUsername(adminUsername)
    if err == nil && existingAdmin != nil {
        log.Println("Admin user already exists")
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
    if err != nil {
        log.Fatalf("Failed to hash password: %v", err)
    }

    admin := models.User{
        Username: adminUsername,
        Password: string(hashedPassword),
        Role:     utils.Admin,
    }

    if err := dao.CreateUser(&admin); err != nil {
        log.Fatalf("Failed to create admin user: %v", err)
    }

    log.Println("Admin user created successfully")
}