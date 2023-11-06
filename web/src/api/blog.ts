import { client } from './client'
import type { Response } from '@/types/index'
import type { Blog, Blogs, FindAllBlogsRequest } from '@/types/blog'

export const FindBlogById = (id: number) => client.get<Response<Blog>>(`/api/v1/blog/${id}`)

export const FindAllBlogs = (params: FindAllBlogsRequest) =>
  client.get<Response<Blogs>>('/api/v1/blog', { params: params })

export const CreateBlog = (blog: Blog) => client.post<Response>('/api/v1/blog', blog)

export const UpdateBlog = (blog: Blog) => client.put<Response>(`/api/v1/blog/${blog.id}`, blog)

export const DeleteBlog = (id: number) => client.delete<Response>(`/api/v1/blog/${id}`)
