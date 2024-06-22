package controllers

import (
    "strconv"
    "net/http"
    "github.com/gin-gonic/gin"
    "books-manage-system/services"
)

// LockUser godoc
// @Summary 锁定用户
// @Description 根据ID锁定用户
// @Tags users
// @id LockUser
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} CommonResponse "User locked successfully"
// @Failure 404 {object} CommonResponse "User not found"
// @Router /users/{id}/lock [patch]
func LockUser(c *gin.Context) {
    id := c.Param("id")
    if err := services.LockUser(id); err != nil {
        c.JSON(http.StatusNotFound, CommonResponse{Message: "User not found"})
        return
    }

    c.JSON(http.StatusOK, CommonResponse{Message: "User locked successfully"})
}

// UnlockUser godoc
// @Summary 解锁用户
// @Description 根据ID解锁用户
// @Tags users
// @id UnlockUser
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} CommonResponse "User locked successfully"
// @Failure 404 {object} CommonResponse "User not found"
// @Router /users/{id}/unlock [patch]
func UnlockUser(c *gin.Context) {
    id := c.Param("id")
    if err := services.UnlockUser(id); err != nil {
        c.JSON(http.StatusNotFound, CommonResponse{Message: "User not found"})
        return
    }

    c.JSON(http.StatusOK, CommonResponse{Message: "User locked successfully"})
}

// GetUsers godoc
// @Summary 获取用户列表
// @Description 获取用户列表并支持分页
// @Tags users
// @id GetUsers
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param page_size query int true "Page size"
// @Success 200 {object} map[string]interface{} "成功返回用户列表"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 404 {object} map[string]interface{} "Records not found"
// @Router /users [get]
func GetUsers(c *gin.Context) {
    // 从查询参数中获取 page，并转换为 int
    pageStr := c.DefaultQuery("page", "1")
    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid page number"})
        return
    }

    // 从查询参数中获取 page_size，并转换为 int
    pageSizeStr := c.DefaultQuery("page_size", "10")
    pageSize, err := strconv.Atoi(pageSizeStr)
    if err != nil || pageSize < 1 {
        c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid page size"})
        return
    }

    // 调用服务获取用户列表
    users, total, err := services.GetUsers(page, pageSize)
    if err != nil {
        c.JSON(http.StatusNotFound, map[string]interface{}{"message": err.Error()})
        return
    }

    // 返回结果
    c.JSON(http.StatusOK, gin.H{
        "data":     users,
        "total":    total,
        "page":     page,
        "page_size": pageSize,
    })
}