import { client } from './client'
import type { Response } from '@/types/type'
import type { LoginForm, LogoutForm, LoginResp } from '@/types/token'

export const LoginReq = (data: LoginForm) => client.post<Response<LoginResp>>('/api/v1/token', data)

export const LogoutReq = (data: LogoutForm) => client.delete<Response<null>>(`/api/v1/token/${data.id}`)
