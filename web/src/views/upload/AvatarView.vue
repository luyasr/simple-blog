<template>
  <div>
    <a-space direction="vertical" :style="{ width: '100%' }">
      <a-upload
        action="/"
        :fileList="file ? [file] : []"
        :show-file-list="false"
        @change="onChange"
        @progress="onProgress"
      >
        <template #upload-button>
          <div
            :class="`arco-upload-list-item${
              file && file.status === 'error' ? ' arco-upload-list-item-error' : ''
            }`"
          >
            <div class="arco-upload-list-picture custom-upload-avatar" v-if="file && file.url">
              <img :src="file.url" />
              <div class="arco-upload-list-picture-mask">
                <IconEdit />
              </div>
              <a-progress
                v-if="file.status === 'uploading' && file.percent < 100"
                :percent="file.percent"
                type="circle"
                size="mini"
                :style="{
                  position: 'absolute',
                  left: '50%',
                  top: '50%',
                  transform: 'translateX(-50%) translateY(-50%)'
                }"
              />
            </div>
            <div class="arco-upload-picture-card" v-else>
              <div class="arco-upload-picture-card-text">
                <IconPlus />
                <div style="margin-top: 10px; font-weight: 600">Upload</div>
              </div>
            </div>
          </div>
        </template>
      </a-upload>
    </a-space>
  </div>
</template>

<script setup lang="ts">
import { IconEdit, IconPlus } from '@arco-design/web-vue/es/icon'
import { ref } from 'vue'

const file = ref()

const onChange = (_: any, currentFile: any) => {
  file.value = {
    ...currentFile
    // url: URL.createObjectURL(currentFile.file),
  }
}
const onProgress = (currentFile: any) => {
  file.value = currentFile
}
</script>

<style scoped lang="less"></style>
