package services

import (
    "books-manage-system/models"
    "books-manage-system/dao"
    "books-manage-system/config"
    "gorm.io/gorm"
    "fmt"
    "time"
)

func CreateBook(book *models.Book) (*models.Book, error) {
    return book, dao.CreateBook(book)
}

func GetBookByID(id uint) (*models.Book, error) {
    return dao.GetBookByID(nil, id)
}

func GetBookList(page int, pageSize int) ([]*models.Book, int64, error) {
    return dao.GetBookList(page, pageSize)
}

func UpdateBook(book *models.Book) (*models.Book, error) {
    book.UpdatedAt = time.Now()
    return book, dao.UpdateBook(nil, book)
}

func DeleteBook(id uint) error {
    return dao.DeleteBook(id)
}

func BorrowBook(userID uint, bookID uint) (models.Borrow, error) {
    var borrow models.Borrow

    // 开始事务
    err := config.DB.Transaction(func(tx *gorm.DB) error {
        // 检查书的库存是否足够
        book, err := dao.GetBookByID(tx, bookID)
        if err != nil {
            return err
        }
        if book.Quantity < 1 {
            return fmt.Errorf("book out of stock")
        }

        // 更新书的库存
        book.Quantity -= 1
        if err := dao.UpdateBook(tx, book); err != nil {
            return err
        }

        // 创建借书记录
        borrow = models.Borrow{
            UserID:    userID,
            BookID:    bookID,
        }
        if err := dao.CreateBorrow(tx, &borrow); err != nil {
            return err
        }

        return nil
    })

    return borrow, err
}

func ReturnBook(borrowID uint) (models.Borrow, error) {
    var borrow models.Borrow

    // 开始事务
    err := config.DB.Transaction(func(tx *gorm.DB) error {
        // 获取借书记录
        borrow, err := dao.GetBorrowByID(tx, borrowID)
        if err != nil {
            return err
        }
        if borrow.ReturnedAt != nil {
            return fmt.Errorf("book already returned")
        }

        // 更新书的库存
        book, err := dao.GetBookByID(tx, borrow.BookID)
        if err != nil {
            return err
        }
        book.Quantity += 1
        if err := dao.UpdateBook(tx, book); err != nil {
            return err
        }

        // 更新借书记录的还书时间
        now := time.Now()
        borrow.ReturnedAt = &now
        if err := dao.UpdateBorrow(tx, borrow); err != nil {
            return err
        }

        return nil
    })

    return borrow, err
}

func GetBorrowedBooks(userID uint, page int, pageSize int) ([]*models.Borrow, int64, error) {
    return dao.GetBorrowedBooks(userID, page, pageSize)
}
