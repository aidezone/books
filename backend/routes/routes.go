package routes

import (
    "github.com/gin-gonic/gin"
    "books-manage-system/controllers"
    "books-manage-system/middleware"
)

func RegisterRoutes(r *gin.Engine) {
    r.Static("/js", "./html/js")
    r.Static("/img", "./html/img")
    r.Static("/assets", "./html/assets")
    r.StaticFile("/", "./html/index.html")
    r.StaticFile("/login", "./html/index.html")
    r.StaticFile("/index.html", "./html/index.html")
    r.StaticFile("/favicon.ico", "./html/favicon.ico")

    v1 := r.Group("/api/v1")
    {

        v1.POST("/register", controllers.Register)
        v1.POST("/login", controllers.Login)

        bookRoutes := v1.Group("/books")
        bookRoutes.Use(middleware.AuthMiddleware())
        {
            bookRoutes.GET("/", controllers.GetBookList)
            bookRoutes.GET("/:id", controllers.GetBook)
            bookRoutes.PATCH("/borrow/:book_id", controllers.BorrowBook)
            bookRoutes.PATCH("/return/:borrow_id", controllers.ReturnBook)
            bookRoutes.Use(middleware.AuthAdminMiddleware())
            {
                bookRoutes.POST("/", controllers.CreateBook)
                bookRoutes.PUT("/:id", controllers.UpdateBook)
                bookRoutes.DELETE("/:id", controllers.DeleteBook)
                bookRoutes.GET("/borrowed/:user_id", controllers.GetBorrowedBooks)
            }
        }

        userRoutes := v1.Group("/users")
        userRoutes.Use(middleware.AuthMiddleware())
        userRoutes.Use(middleware.AuthAdminMiddleware())
        {
            userRoutes.GET("/", controllers.GetUsers)
            userRoutes.PATCH("/:id/lock", controllers.LockUser)
            userRoutes.PATCH("/:id/unlock", controllers.UnlockUser)
        }
    }
}
