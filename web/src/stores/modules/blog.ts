import { FindAllBlogs } from '@/api/blog'
import type { Blogs, FindAllBlogsRequest } from '@/types/blog'
import { defineStore } from 'pinia'
import dayjs from 'dayjs'

export const useBlogStore = defineStore({
  id: 'blog',
  state: () => ({
    blogs: {} as Blogs
  }),
  actions: {
    async findAllBlogs(params: FindAllBlogsRequest) {
      const resp = await FindAllBlogs(params)
      if (resp.code == 200) {
        this.blogs = resp.data
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
