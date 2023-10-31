import { FindAllBlogs, UpdateContent } from '@/api/blog'
import type { Blogs, FindAllBlogsRequest, UpdateBlogRequest } from '@/types/blog'
import { defineStore } from 'pinia'
import dayjs from 'dayjs'

export const useBlogStore = defineStore({
  id: 'blog',
  state: () => ({
    blogs: {} as Blogs,
    updateBlog: {} as UpdateBlogRequest
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
    },
    async updateText(text: string) {
      const data: UpdateBlogRequest = {
        id: this.updateBlog.id,
        content: text
      }
      const resp = await UpdateContent(data)
      if (resp.code == 200) {
        return 'ok'
      } else {
        return Promise.reject(new Error(resp.message))
      }
    },
    setUpdateBlog(updateBlog: UpdateBlogRequest) {
      this.updateBlog = updateBlog
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
