import axios, { type AxiosRequestConfig, type AxiosResponse } from 'axios'
import { useUserStore } from '@/stores/modules/user'
import { Message } from '@arco-design/web-vue'

const instance = axios.create({
  baseURL: '',
  timeout: 5000,
  headers: { 'Content-Type': 'application/json' }
})

// 请求拦截器
instance.interceptors.request.use((config) => {
  const userStore = useUserStore()

  if (userStore.access_token) {
    config.headers.access_token = userStore.access_token
  }
  return config
})

// 响应拦截器
instance.interceptors.response.use(
  (response: AxiosResponse) => {
    return response.data
  },
  (error) => {
    if (error.response) {
      // 请求已发送，服务器返回错误响应
      const { status, data } = error.response
      switch (status) {
        case 400:
          Message.error('错误的请求')
          break
        case 401:
          Message.error('未授权的访问')
          break
        case 403:
          Message.error('禁止访问')
          break
        case 404:
          Message.error('请求的资源不存在')
          break
        case 500:
          Message.error('服务器内部错误')
          break
        default:
          Message.error(`请求错误，状态码：${status}`)
      }
      return Promise.reject(data)
    } else if (error.request) {
      // 请求已发送，但没有收到响应
      Message.error('请求超时，请检查网络连接')
      return Promise.reject(error)
    } else {
      // 请求未发送，发生了错误
      Message.error('请求错误，请检查网络连接')
      return Promise.reject(error)
    }
  }
)

export const client = {
  get<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
    return instance.request({
      url: url,
      method: 'get',
      ...config
    })
  },
  post<T = any>(url: string, data?: object, config?: AxiosRequestConfig): Promise<T> {
    return instance.request({
      url: url,
      method: 'post',
      data: data,
      ...config
    })
  },
  put<T = any>(url: string, data?: object, config?: AxiosRequestConfig): Promise<T> {
    return instance.request({
      url: url,
      method: 'put',
      data: data,
      ...config
    })
  },
  delete<T = any>(url: string, data?: object, config?: AxiosRequestConfig): Promise<T> {
    return instance.request({
      url: url,
      method: 'delete',
      data: data,
      ...config
    })
  }
}
