export interface UserResponse {
  id: string
  create_at: number
  update_at: number
  delete_at: number
  username: string
  password: string
  avatar: string
  role: string
}

export interface LoginRequest {
  username: string
  password: string
}

export interface LogoutRequest {
  access_token: string | null
  refresh_token: string | null
}

export interface LoginResponse {
  id: number
  create_at: number
  update_at: number
  user_id: number
  username: string
  access_token: string
  access_token_expired_at: number
  refresh_token: string
  refresh_token_expired_at: number
}
