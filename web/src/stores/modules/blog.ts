import { FindBlogById, FindAllBlogs, CreateBlog, UpdateBlog, DeleteBlog } from '@/api/blog'
import type { Blog, Blogs, FindAllBlogsRequest } from '@/types/blog'
import { defineStore } from 'pinia'
import dayjs from 'dayjs'

export const useBlogStore = defineStore({
  id: 'blog',
  state: () => ({
    blog: {} as Blog,
    blogs: {} as Blogs
  }),
  actions: {
    async findBlogById(id: number) {
      const resp = await FindBlogById(id)
      if (resp.code == 200) {
        this.blog = resp.data
        return 'ok'
      } else {
        return Promise.reject(new Error(resp.message))
      }
    },
    async findAllBlogs(params: FindAllBlogsRequest) {
      const resp = await FindAllBlogs(params)
      if (resp.code == 200) {
        this.blogs = resp.data
        return 'ok'
      } else {
        return Promise.reject(new Error(resp.message))
      }
    },
    async createBlog(blog: Blog) {
      const resp = await CreateBlog(blog)
      if (resp.code == 200) {
        return 'ok'
      } else {
        return Promise.reject(new Error(resp.message))
      }
    },
    async updateBlog(blog: Blog) {
      const resp = await UpdateBlog(blog)
      if (resp.code == 200) {
        return 'ok'
      } else {
        return Promise.reject(new Error(resp.message))
      }
    },
    async deleteBlog(id: number) {
      const resp = await DeleteBlog(id)
      if (resp.code == 200) {
        return 'ok'
      } else {
        return Promise.reject(new Error(resp.message))
      }
    }
  },
  getters: {
    getBlogs(): Blob[] {
      return this.blogs.items.map((blog: any) => {
        return {
          ...blog,
          create_at: dayjs.unix(blog.create_at).format('YYYY-MM-DD HH:mm:ss'),
          update_at: dayjs.unix(blog.update_at).format('YYYY-MM-DD HH:mm:ss'),
          audit_at: dayjs.unix(blog.audit_at).format('YYYY-MM-DD HH:mm:ss')
        }
      })
    },
    getTotal(): number {
      return this.blogs.total
    }
  }
})
