package controllers

import (
    "fmt"
	"strconv"
    "net/http"
    "github.com/gin-gonic/gin"
    // "gorm.io/gorm"
    // "github.com/dgrijalva/jwt-go"
    "books-manage-system/utils"
    "books-manage-system/models"
    "books-manage-system/services"
)

// 定义响应结构体
type BookResponse *models.BookInfo
type BorrowResponse *models.Borrow

type BooksResponse struct {
    Books []*models.Book `json:"books"`
    Total int64              `json:"total"`
}
type BorrowListResponse struct {
    BorrowList []*models.Borrow `json:"borrow_list"`
    Total int64              `json:"total"`
}
// GetBook godoc
// @Summary 获取图书
// @Description 根据ID获取图书
// @Tags books
// @id GetBook
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} BookResponse
// @Failure 404 {object} CommonResponse
// @Router /books/{id} [get]
func GetBook(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusNotFound, CommonResponse{Message: "Param invalid"})
        return
    }
    book, err := services.GetBookByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, CommonResponse{Message: "Book not found"})
        return
    }

    c.JSON(http.StatusOK, book.BookInfo)
}

// GetBookList godoc
// @Summary 查询图书
// @Description 获取图书列表
// @Tags books
// @id GetBookList
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param keyword query string true "关键词"
// @Param search_type query int true "检索类型"
// @Param page query int false "页码"
// @Param page_size query int false "分页大小"
// @Success 200 {object} BooksResponse
// @Failure 404 {object} CommonResponse
// @Router /books [get]
func GetBookList(c *gin.Context) {
    page, _ := strconv.Atoi(c.Param("page"))
    pageSize, _ := strconv.Atoi(c.Param("page_size"))

    if page < 1 {
        page = 1
    }

    if pageSize > 200 {
        pageSize = 200
    }

    if pageSize < 10 {
        pageSize = 10
    }
    
    bookList, total, err := services.GetBookList(page, pageSize)
    if err != nil {
        c.JSON(http.StatusNotFound, CommonResponse{Message: "Book not found"})
        return
    }
    // var bookResp []*models.BookInfo
    // for _, book := range bookList {
    //     bookResp = append(bookResp, &book.BookInfo)
    // }
    c.JSON(http.StatusOK, BooksResponse{
        Books: bookList,
        Total: total,
    })
}

// CreateBook godoc
// @Summary 创建图书
// @Description 创建新图书
// @Tags books
// @id CreateBook
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param book body models.BookInfo true "Book"
// @Success 200 {object} BookResponse
// @Failure 400 {object} CommonResponse
// @Router /books [post]
func CreateBook(c *gin.Context) {
    var book models.BookInfo
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, CommonResponse{Message: fmt.Sprintf("Invalid input, err: %v", err)})
        return
    }
    bookModel := &models.Book{
        BookInfo: book,
    }

    createdBook, err := services.CreateBook(bookModel)
    if err != nil {
        c.JSON(http.StatusInternalServerError, CommonResponse{Message: "Failed to create book"})
        return
    }

    c.JSON(http.StatusCreated, createdBook)
}

// UpdateBook godoc
// @Summary 更新图书
// @Description 根据ID更新图书
// @Tags books
// @id UpdateBook
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.BookInfo true "Book"
// @Success 200 {object} BookResponse
// @Failure 400 {object} CommonResponse
// @Failure 404 {object} CommonResponse
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusNotFound, CommonResponse{Message: "Param invalid"})
        return
    }

    var book models.BookInfo
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, CommonResponse{Message: "Invalid input"})
        return
    }

    bookModel, err := services.GetBookByID(uint(id))
    bookModel.BookInfo = book
    updatedBook, err := services.UpdateBook(bookModel)
    if err != nil {
        c.JSON(http.StatusNotFound, CommonResponse{Message: fmt.Sprintf("Book not found, bookinfo: %+v, err: %+v", bookModel, err)})
        return
    }

    c.JSON(http.StatusOK, updatedBook)
}

// DeleteBook godoc
// @Summary 删除图书
// @Description 根据ID删除图书
// @Tags books
// @id DeleteBook
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} CommonResponse
// @Failure 404 {object} CommonResponse
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusNotFound, CommonResponse{Message: "Param invalid"})
        return
    }

    if err := services.DeleteBook(uint(id)); err != nil {
        c.JSON(http.StatusNotFound, CommonResponse{Message: "Book not found"})
        return
    }

    c.JSON(http.StatusOK, CommonResponse{Message: "Book deleted successfully"})
}

// BorrowBook godoc
// @Summary 借书
// @Description 用户借书
// @Tags books
// @id BorrowBook
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param book_id path int true "Book ID"
// @Success 200 {object} BorrowResponse
// @Failure 400 {object} CommonResponse "Invalid input"
// @Failure 404 {object} CommonResponse "Book not found or out of stock"
// @Router /books/borrow/{book_id} [patch]
func BorrowBook(c *gin.Context) {
    user, ok := c.Get("user")
    if !ok {
        c.JSON(http.StatusUnauthorized, CommonResponse{Message: "Invalid context"})
        c.Abort()
    }
    userID := user.(*utils.Claims).UserID
    bookID, err := strconv.ParseUint(c.Param("book_id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, CommonResponse{Message: "Invalid book ID"})
        return
    }

    borrow, err := services.BorrowBook(uint(userID), uint(bookID))
    if err != nil {
        c.JSON(http.StatusNotFound, CommonResponse{Message: err.Error()})
        return
    }

    c.JSON(http.StatusOK, borrow)
}

// ReturnBook godoc
// @Summary 还书
// @Description 用户还书
// @Tags books
// @id ReturnBook
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param borrow_id path int true "Borrow ID"
// @Success 200 {object} BorrowResponse
// @Failure 400 {object} CommonResponse "Invalid input"
// @Failure 404 {object} CommonResponse "Borrow record not found"
// @Router /books/return/{borrow_id} [patch]
func ReturnBook(c *gin.Context) {
    borrowID, err := strconv.ParseUint(c.Param("borrow_id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, CommonResponse{Message: "Invalid borrow ID"})
        return
    }

    borrow, err := services.ReturnBook(uint(borrowID))
    if err != nil {
        c.JSON(http.StatusNotFound, CommonResponse{Message: err.Error()})
        return
    }

    c.JSON(http.StatusOK, borrow)
}

// GetBorrowedBooks godoc
// @Summary 获取已借图书列表
// @Description 获取用户已借图书列表并支持分页
// @Tags books
// @id GetBorrowedBooks
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param page query int true "Page number"
// @Param page_size query int true "Page size"
// @Success 200 {object} BorrowListResponse "成功返回借书记录列表"
// @Failure 400 {object} CommonResponse "Invalid input"
// @Failure 404 {object} CommonResponse "Records not found"
// @Router /books/borrowed/{user_id} [get]
func GetBorrowedBooks(c *gin.Context) {
    userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, CommonResponse{Message: "Invalid user ID"})
        return
    }

    page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
    if err != nil {
        c.JSON(http.StatusBadRequest, CommonResponse{Message: "Invalid page number"})
        return
    }

    pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
    if err != nil {
        c.JSON(http.StatusBadRequest, CommonResponse{Message: "Invalid page size"})
        return
    }

    if page < 1 {
        page = 1
    }

    if pageSize > 200 {
        pageSize = 200
    }

    if pageSize < 10 {
        pageSize = 10
    }

    borrows, total, err := services.GetBorrowedBooks(uint(userID), page, pageSize)
    if err != nil {
        c.JSON(http.StatusNotFound, CommonResponse{Message: err.Error()})
        return
    }

    c.JSON(http.StatusOK, BorrowListResponse{
        BorrowList:  borrows,
        Total: total,
    })
}