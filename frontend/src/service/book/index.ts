import request from '../../fetch/axios/instance'
import { API } from '@/fetch/api'
import fetch from '@/fetch/index'

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
export const GetBookList = async (params: any) => {
  try {
    const response = await fetch(API.books.GetBookList, {
      body: params
    })
    return response
  } catch (error) {
    return Promise.reject(error)
  }
}
