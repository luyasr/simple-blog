<script setup lang="ts">
import router from '@/router'
import Menu from './MenuView.vue'
// 获取父组件传递过来的路由数组
defineProps(['menuList'])

const handleClick = (key: string) => {
  router.push(key)
}
</script>

<template>
  <div class="menu">
    <a-menu :style="{ width: '100%' }" @menu-item-click="handleClick">
      <template v-for="item in menuList" :key="item.path">
        <!-- 路由隐藏 -->
        <template v-if="!item.meta.hideInMenu">
          <!-- 没有子路由 -->
          <a-menu-item v-if="!item.children" :key="item.path">
            <template #icon>
              <component :is="item.meta.icon"></component>
            </template>
            <span>{{ item.meta.title }}</span>
          </a-menu-item>
        </template>

        <!-- 有一个子路由 -->
        <a-menu-item v-if="item.children && item.children.length == 1" :key="item.children[0].path">
          <template #icon>
            <component :is="item.children[0].meta.icon"></component>
          </template>
          <span>{{ item.children[0].meta.title }}</span>
        </a-menu-item>

        <!-- 有子路由则递归 -->
        <a-sub-menu v-if="item.children && item.children.length > 1" :key="item.path">
          <template #icon>
            <component :is="item.meta.icon"></component>
          </template>
          <template #title><span>{{ item.meta.title }}</span></template>
          <!-- 递归组件 https://cn.vuejs.org/api/sfc-script-setup.html#recursive-components -->
          <Menu :menuList="item.children" />
        </a-sub-menu>
      </template>
    </a-menu>
  </div>
</template>

<style scoped lang="less">
.menu {
  height: calc(100% - @base-menu-logo-height);
}
</style>
