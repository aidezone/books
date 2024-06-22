declare namespace Model.books {
    export interface ControllersBookResponse {
        author: string;
        category: string;
        quantity: number;
        title: string;
    }
    export interface ControllersBooksResponse {
        books?: ModelsBook[];
        total?: number;
    }
    export interface ControllersBorrowListResponse {
        borrow_list?: ModelsBorrow[];
        total?: number;
    }
    export interface ControllersBorrowResponse {
        book_id?: number;
        createdAt?: string;
        deletedAt?: GormDeletedAt;
        id?: number;
        returned_at?: string;
        updatedAt?: string;
        user_id?: number;
    }
    export interface ControllersCommonResponse {
        code?: number;
        message?: string;
    }
    export interface ControllersLoginResponse {
        token?: string;
    }
    export interface ControllersUserRequest {
        password: string;
        username: string;
    }
    export interface GormDeletedAt {
        time?: string;
        /**
         * Valid is true if Time is not NULL
         */
        valid?: boolean;
    }
    export interface ModelsBook {
        author: string;
        category: string;
        createdAt?: string;
        deletedAt?: GormDeletedAt;
        id?: number;
        quantity: number;
        title: string;
        updatedAt?: string;
    }
    export interface ModelsBookInfo {
        author: string;
        category: string;
        quantity: number;
        title: string;
    }
    export interface ModelsBorrow {
        book_id?: number;
        createdAt?: string;
        deletedAt?: GormDeletedAt;
        id?: number;
        returned_at?: string;
        updatedAt?: string;
        user_id?: number;
    }
}
