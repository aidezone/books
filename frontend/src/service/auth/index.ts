import request from "../../fetch/axios/instance";
import { API } from "@/fetch/api";
import fetch from "@/fetch/index";

/**
 * 用户通过用户名和密码登录
 * 
 * @param {Model.books.ControllersUserRequest} request - [body] payload (required)
 * @return {Model.books.ControllersLoginResponse} Model.books.ControllersLoginResponse
 * @summary POST /login
 */
export const accountLogin = async (params: any) => {
  try {
    const response = await fetch(API.books.Login, {
      body: params
    });
    return response;
  } catch (error) {
    return Promise.reject(error);
  }
};

/**
 * 通过用户名和密码注册新用户
 * 
 * @param {Model.books.ControllersUserRequest} request - [body] payload (required)
 * @return {Model.books.ControllersUserRequest} Model.books.ControllersUserRequest
 * @summary POST /register
 */
export const accountRegister = async (params: any) => {
  try {
    const response = await fetch(API.books.Register, {
      body: params
    });
    return response;
  } catch (error) {
    return Promise.reject(error);
  }
};