import { client } from './client'
import type { Response } from '@/types/index'
import type { LoginRequest, LogoutRequest, LoginResponse, UserResponse } from '@/types/user'

export const Login = (data: LoginRequest) =>
  client.post<Response<LoginResponse>>('/api/v1/token', data)

export const Logout = (id: number, data: LogoutRequest) =>
  client.delete<Response<null>>(`/api/v1/token/${id}`, data)

export const UserInfo = (id: number) => client.get<Response<UserResponse>>(`/api/v1/user/${id}`)
