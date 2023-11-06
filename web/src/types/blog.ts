export interface FindAllBlogsRequest {
  // 博客状态
  status?: number
  // 博客分页大小
  page_size?: number
  // 博客分页页码
  page_number?: number
  // 博客查询关键字
  keywords?: string
  // 博客用户数组
  usernames?: string[]
}

export interface Blog {
  // 博客ID
  id: number
  // 博客作者
  author: string
  // 博客标题
  title: string
  // 博客摘要
  summary: string
  // 博客内容
  content: string
  // 博客标签
  tags: Map<string, string>
  // 博客创建用户
  create_by: string
  // 博客状态
  status: number
  // 博客审核人
  audit_at: number
  // 博客审核状态
  audit_status: number
  // 博客创建时间
  create_at: number
  // 博客更新时间
  update_at: number
}

export interface Blogs {
  // 博客总数
  total: number
  items: Blog[]
}

export enum BlogStatus {
  // 博客状态：草稿
  DRAFT = 0,
  // 博客状态：发布
  PUBLISH = 1,
  // 博客状态：删除
  DELETE = 2
}
