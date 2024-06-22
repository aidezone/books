/* This file was auto-generated. Don't modify this manually. */

export namespace API {
  // books
  export enum books {
    /**
     * 获取图书列表
     *
     * @param {string} keyword - [query] 关键词 (required)
     * @param {integer} search_type - [query] 检索类型 (required)
     * @param {integer} page - [query] 页码
     * @param {integer} page_size - [query] 分页大小
     * @return {Model.books.ControllersBooksResponse} Model.books.ControllersBooksResponse
     * @summary GET /books
     */
    GetBookList = 'books:GetBookList',

    /**
     * 创建新图书
     *
     * @param {Model.books.ModelsBookInfo} book - [body] Book (required)
     * @return {Model.books.ControllersBookResponse} Model.books.ControllersBookResponse
     * @summary POST /books
     */
    CreateBook = 'books:CreateBook',

    /**
     * 用户借书
     *
     * @param {integer} book_id - [path] Book ID (required)
     * @return {Model.books.ControllersBorrowResponse} Model.books.ControllersBorrowResponse
     * @summary PATCH /books/borrow/{book_id}
     */
    BorrowBook = 'books:BorrowBook',

    /**
     * 获取用户已借图书列表并支持分页
     *
     * @param {integer} user_id - [path] User ID (required)
     * @param {integer} page - [query] Page number (required)
     * @param {integer} page_size - [query] Page size (required)
     * @return {Model.books.ControllersBorrowListResponse} Model.books.ControllersBorrowListResponse
     * @summary GET /books/borrowed/{user_id}
     */
    GetBorrowedBooks = 'books:GetBorrowedBooks',

    /**
     * 用户还书
     *
     * @param {integer} borrow_id - [path] Borrow ID (required)
     * @return {Model.books.ControllersBorrowResponse} Model.books.ControllersBorrowResponse
     * @summary PATCH /books/return/{borrow_id}
     */
    ReturnBook = 'books:ReturnBook',

    /**
     * 根据ID获取图书
     *
     * @param {integer} id - [path] Book ID (required)
     * @return {Model.books.ControllersBookResponse} Model.books.ControllersBookResponse
     * @summary GET /books/{id}
     */
    GetBook = 'books:GetBook',

    /**
     * 根据ID更新图书
     *
     * @param {integer} id - [path] Book ID (required)
     * @param {Model.books.ModelsBookInfo} book - [body] Book (required)
     * @return {Model.books.ControllersBookResponse} Model.books.ControllersBookResponse
     * @summary PUT /books/{id}
     */
    UpdateBook = 'books:UpdateBook',

    /**
     * 根据ID删除图书
     *
     * @param {integer} id - [path] Book ID (required)
     * @return {Model.books.ControllersCommonResponse} Model.books.ControllersCommonResponse
     * @summary DELETE /books/{id}
     */
    DeleteBook = 'books:DeleteBook',

    /**
     * 用户通过用户名和密码登录
     *
     * @param {Model.books.ControllersUserRequest} request - [body] payload (required)
     * @return {Model.books.ControllersLoginResponse} Model.books.ControllersLoginResponse
     * @summary POST /login
     */
    Login = 'books:Login',

    /**
     * 通过用户名和密码注册新用户
     *
     * @param {Model.books.ControllersUserRequest} request - [body] payload (required)
     * @return {Model.books.ControllersUserRequest} Model.books.ControllersUserRequest
     * @summary POST /register
     */
    Register = 'books:Register',

    /**
     * 获取用户列表并支持分页
     *
     * @param {integer} page - [query] Page number (required)
     * @param {integer} page_size - [query] Page size (required)
     * @return {object} object
     * @summary GET /users
     */
    GetUsers = 'books:GetUsers',

    /**
     * 根据ID锁定用户
     *
     * @param {integer} id - [path] User ID (required)
     * @return {Model.books.ControllersCommonResponse} Model.books.ControllersCommonResponse
     * @summary PATCH /users/{id}/lock
     */
    LockUser = 'books:LockUser',

    /**
     * 根据ID解锁用户
     *
     * @param {integer} id - [path] User ID (required)
     * @return {Model.books.ControllersCommonResponse} Model.books.ControllersCommonResponse
     * @summary PATCH /users/{id}/unlock
     */
    UnlockUser = 'books:UnlockUser'
  }
}
