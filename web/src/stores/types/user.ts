export interface UserState {
  access_token: string | null
  refresh_token: string | null
  isAuthenticated: boolean
  user_id: number
  username: string
  avatar: string
}
