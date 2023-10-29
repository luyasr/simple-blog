<script setup lang="ts">
import settings from '@/config/settings'
</script>

<template>
  <div class="logo" v-if="!settings.hideInLogo">
    <img :src="settings.logo" alt="" />
    <span>{{ settings.title }}</span>
  </div>
</template>

<style scoped lang="less">
.logo {
  display: flex;
  height: 32px;
  margin: 8px 15px;
  justify-content: flex-start;
  align-items: center;
  background: rgba(255, 255, 255, 0.2);
  overflow: hidden;
  /* 隐藏溢出部分 */

  img {
    display: flex;
    width: 20px;
    height: 20px;
    align-items: center;
  }

  span {
    font-size: @base-menu-title-font-size;
    margin-left: 15px;
    /* 隐藏溢出部分 */
    overflow: hidden;
    /* 防止文本换行 */
    white-space: nowrap;
    /* 截断文本内容 */
    text-overflow: ellipsis;
    font-weight: normal;
    box-sizing: border-box;
    /* 调整菜单项的宽度 */
    width: 100%;
  }
}
</style>
