<template>
  <MdEditor v-model="text" @onSave="onSave" />
</template>

<script setup lang="ts">
import { onBeforeMount, ref } from 'vue'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import { useBlogStore } from '@/stores/modules/blog'
import { Message } from '@arco-design/web-vue'
import type { UpdateBlogRequest } from '@/types/blog'

const text = ref('')
const blogStore = useBlogStore()

// 挂载前，将 store 中的 text 赋值给 text
onBeforeMount(() => {
  text.value = blogStore.updateBlog.content
})

const onSave = async (text: string) => {
  try {
    // store 中的 text 与当前编辑器中的 text 不一致时，更新
    if (blogStore.updateBlog.content != text) {
      await blogStore.updateText(text)

      // 更新 store 中的 updateBlog
      let updateBlog: UpdateBlogRequest = {
        id: blogStore.updateBlog.id,
        content: text
      }
      blogStore.setUpdateBlog(updateBlog)
      Message.info('保存成功')
    }
  } catch (error) {
    Message.error(`${error}`)
  }
}
</script>

<style scoped></style>
