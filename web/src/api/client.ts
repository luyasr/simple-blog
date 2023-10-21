import axios from 'axios'

export const client = axios.create({
  baseURL: '',
  timeout: 5000,
  headers: { 'Content-Type': 'application/json' }
})
