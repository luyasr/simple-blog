<template>
  <div>
    <div class="blog-tab">
      <a-button type="primary" @click="router.push({ name: 'blogEdit' })">添加</a-button>
      <a-input-search
        v-model="params.keywords"
        :style="{ width: '320px' }"
        :loading="searchLoading"
        placeholder="请输入搜索关键词"
        @press-enter="onSearch"
      />
    </div>
    <div>
      <a-table :columns="columns" :data="blogs" :loading="loading" :pagination="false">
        <template #status="{ record }">
          <a-tag v-if="record.status == BlogStatus.DRAFT" color="gold" bordered>草稿</a-tag>
          <a-tag v-if="record.status == BlogStatus.PUBLISH" color="green" bordered>发布</a-tag>
        </template>
        <template #optional="{ record }">
          <a-button
            type="primary"
            @click="router.push({ name: 'blogEdit', query: { id: record.id } })"
          >
            编辑
          </a-button>
          <a-popconfirm content="您确定要删除吗?" type="error" @ok="deleteBlog(record.id)">
            <a-button status="danger">删除</a-button>
          </a-popconfirm>
        </template>
      </a-table>
    </div>
    <div class="pagination">
      <a-pagination
        :total="total"
        show-total
        show-jumper
        show-page-size
        :hide-on-single-page="true"
        @change="onPageNumberChange"
        @page-size-change="onPageSizeChange"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useBlogStore } from '@/stores/modules/blog'
import { onBeforeMount } from 'vue'
import { BlogStatus } from '@/types/blog'
import type { FindAllBlogsRequest } from '@/types/blog'
import router from '@/router'

// 查询博客列表参数
const params = reactive({} as FindAllBlogsRequest)
// table 列
const columns = [
  {
    title: '作者',
    dataIndex: 'author'
  },
  {
    title: '标题',
    dataIndex: 'title'
  },
  {
    title: '创建时间',
    dataIndex: 'create_at'
  },
  {
    title: '文章状态',
    dataIndex: 'status',
    slotName: 'status'
  },
  {
    title: '操作',
    slotName: 'optional'
  }
]
// 博客列表
let blogs = reactive([{}])
// 博客总数
const total = ref(0)
// 博客列表加载状态
const loading = ref(false)
// 搜索加载状态
const searchLoading = ref(false)
// 博客 store
let blogStore = useBlogStore()

const findAllBlogs = async (params: FindAllBlogsRequest) => {
  try {
    loading.value = true
    await blogStore.findAllBlogs(params)
    blogs = blogStore.getBlogs
    total.value = blogStore.getTotal
  } finally {
    loading.value = false
  }
}

// 搜索
const onSearch = async () => {
  try {
    searchLoading.value = true
    params.page_number = 1
    await findAllBlogs(params)
  } finally {
    searchLoading.value = false
  }
}

// 分页大小改变
const onPageSizeChange = (pageSize: number) => {
  params.page_size = pageSize
  params.page_number = 1
  findAllBlogs(params)
}

// 页码改变
const onPageNumberChange = (pageNumber: number) => {
  params.page_number = pageNumber
  findAllBlogs(params)
}

// 删除博客
const deleteBlog = async (id: number) => {
  try {
    await blogStore.deleteBlog(id)
    await findAllBlogs(params)
  } catch (error) {
    console.log(error)
  }
}

// 页面加载前请求博客列表
onBeforeMount(async () => {
  await findAllBlogs(params)
})
</script>

<style scoped lang="less">
.blog-tab {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}
.pagination {
  display: flex;
  justify-content: flex-end;
  margin-top: 20px;
}
</style>
