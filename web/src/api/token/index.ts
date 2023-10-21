import { client } from '../client'
import type { Response } from '../type'
import type { LoginForm, LogoutForm, LoginResp } from './type'

export const LoginReq = (data: LoginForm) =>
  client.post<any, Response<LoginResp>>('/api/v1/token', data)

export const LogoutReq = (data: LogoutForm) => client.delete(`/api/v1/token/${data.id}`)
