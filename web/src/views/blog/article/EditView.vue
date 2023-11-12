<template>
  <div>
    <a-page-header :title="title" @back="router.go(-1)" />
    <div :style="{ padding: '20px' }">
      <a-form
        :model="blog"
        layout="vertical"
        auto-label-width
        label-align="left"
        @submit-success="handleSubmit"
        :rules="rules"
      >
        <a-form-item field="title" label="标题" :hide-asterisk="true" :style="{ width: '600px' }">
          <a-input v-model="blog.title" placeholder="请输入标题" />
        </a-form-item>
        <a-form-item field="summary" label="摘要" :hide-asterisk="true" :style="{ width: '600px' }">
          <a-input v-model="blog.summary" placeholder="请输入摘要" />
        </a-form-item>
        <a-form-item field="content" label="内容" :hide-asterisk="true">
          <MdEditor v-model="blog.content" @onSave="handleSubmit" />
        </a-form-item>
        <a-form-item>
          <a-button>取消</a-button>
          <a-button type="primary" html-type="submit">保存</a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onBeforeMount, onMounted, onUnmounted } from 'vue'
import router from '@/router'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import type { Blog } from '@/types/blog'
import { useBlogStore } from '@/stores/modules/blog'
import { Message } from '@arco-design/web-vue'

const title = ref('创建文章')
const blog = reactive({
  title: '',
  summary: '',
  content: ''
} as Blog)
const blogStore = useBlogStore()

const rules = {
  title: [{ required: true, message: '标题是必填项' }],
  summary: [{ required: true, message: '摘要是必填项' }],
  content: [{ required: true, message: '内容是必填项' }]
}

const handleSubmit = async () => {
  // 如果存在id，则是更新, 否则是创建
  try {
    if (
      blog.title === blogStore.blog.title &&
      blog.summary === blogStore.blog.summary &&
      blog.content === blogStore.blog.content
    ) {
      Message.info('未修改')
      return
    }

    if (blog.id) {
      await blogStore.updateBlog(blog)
      Message.info('保存成功')
    } else {
      await blogStore.createBlog(blog)
      Message.info('创建成功')
    }
  } catch (error) {
    console.log(error)
  }
}

const findBlogById = async () => {
  const blogId = router.currentRoute.value.query.id
  if (blogId) {
    title.value = '编辑文章'
    try {
      await blogStore.findBlogById(Number(blogId))
      blog.id = blogStore.blog.id
      blog.title = blogStore.blog.title
      blog.summary = blogStore.blog.summary
      blog.content = blogStore.blog.content
    } catch (error) {
      console.log(error)
    }
  }
}

// 在这个修改后的代码中，我在 onMounted 钩子中添加了一个键盘事件监听器，当按下 Ctrl+S 时，调用 handleSubmit 函数。
// 注意，我使用了 event.preventDefault() 来阻止浏览器的默认 Ctrl+S 行为（通常是保存网页）。
// 然后，在 onUnmounted 钩子中，我移除了这个监听器，以避免内存泄漏
onMounted(() => {
  const handleKeyDown = (e: KeyboardEvent) => {
    if (e.ctrlKey && e.key === 's') {
      e.preventDefault()
      handleSubmit()
    }
  }

  window.addEventListener('keydown', handleKeyDown)

  onUnmounted(() => {
    window.removeEventListener('keydown', handleKeyDown)
  })
})

// 挂载前，将 store 中的 blog 赋值给 blog
onBeforeMount(() => {
  findBlogById()
})
</script>

<style scoped lang="less">
.blog-page-header {
  display: flex;
  justify-content: space-between;
}
</style>
