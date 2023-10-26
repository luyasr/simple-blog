import { client } from './client'
import type { Response } from '@/types/index'
import type { LoginForm, LogoutForm, LoginResp, UserResp } from '@/types/user'

export const LoginReq = (data: LoginForm) => client.post<Response<LoginResp>>('/api/v1/token', data)

export const LogoutReq = (data: LogoutForm) =>
    client.delete<Response<null>>(`/api/v1/token/${data.id}`)

export const UserInfoReq = (id: number) => client.get<Response<UserResp>>(`/api/v1/user/${id}`)
