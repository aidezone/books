import axios from 'axios'
import defaultConfigs from './config'

// 创建 axios 实例
const instance = axios.create({
  ...defaultConfigs,
  // baseURL: 'http://localhost:8080/api', // 替换为你的 Swagger API 基础 URL
  headers: {
    'Content-Type': 'application/json'
  }
})

// 添加请求拦截器，如果有需要在请求中添加 Authorization 头
instance.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

export default instance
