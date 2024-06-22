package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "books-manage-system/utils"
    // "books-manage-system/models"
    "books-manage-system/services"
)

type UserRequest struct {
    Username  string `json:"username" binding:"required"`
    Password  string `json:"password" binding:"required"`
}

// 定义响应结构体
type LoginResponse struct {
    Token string `json:"token"`
}


// Register godoc
// @Summary 注册新用户
// @Description 通过用户名和密码注册新用户
// @Tags auth
// @id Register
// @Accept json
// @Produce json
// @Param request body UserRequest true "payload"
// @Success 200 {object} UserRequest
// @Failure 400 {object} CommonResponse
// @Router /register [post]
func Register(c *gin.Context) {
    var user UserRequest
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, CommonResponse{Code: utils.ERR_PARAM, Message: "Invalid input"})
        return
    }
    err := services.Register(user.Username, user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, CommonResponse{Code: utils.ERR_USER_REGISER, Message: "User register failed!"})
        return
    }

    c.JSON(http.StatusOK, CommonResponse{Code: utils.SUCCESS, Message: "SUCCESS"})
}

// Login godoc
// @Summary 用户登录
// @Description 用户通过用户名和密码登录
// @Tags auth
// @id Login
// @Accept json
// @Produce json
// @Param request body UserRequest true "payload"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} CommonResponse
// @Router /login [post]
func Login(c *gin.Context) {
    var credentials UserRequest
    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, CommonResponse{Message: "Invalid input"})
        return
    }

    user, err := services.Login(credentials.Username, credentials.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, CommonResponse{Message: "Authentication failed"})
        return
    }

    token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
        c.JSON(http.StatusUnauthorized, CommonResponse{Message: "GenerateToken failed"})
        return
    }

    c.JSON(http.StatusOK, LoginResponse{
        Token: token,
    })
}
