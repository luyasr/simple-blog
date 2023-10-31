import { client } from './client'
import type { Response } from '@/types/index'
import type { Blogs, FindAllBlogsRequest, UpdateBlogRequest } from '@/types/blog'

export const FindAllBlogs = (params: FindAllBlogsRequest) =>
  client.get<Response<Blogs>>('/api/v1/blog', { params: params })

export const UpdateContent = (data: UpdateBlogRequest) =>
  client.put<Response>(`/api/v1/blog/${data.id}`, data)
