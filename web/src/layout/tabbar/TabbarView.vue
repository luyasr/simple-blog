<template>
  <div class="layout-breadcrumb">
    <a-space direction="vertical">
      <a-breadcrumb :style="{ margin: '16px 0' }">
        <template v-for="item in $route.matched" :key="item.path">
          <a-breadcrumb-item v-if="item.meta.title">
            <template #separator>
              <span>></span>
            </template>
            <!-- 判断当前组件是否为一级父组件, 如果是就不能跳转 -->
            <router-link v-if="isTopLevelComponent(item.path)" to="">
              {{ item.meta.title }}
            </router-link>
            <router-link v-else :to="item.path">
              {{ item.meta.title }}
            </router-link>
          </a-breadcrumb-item>
        </template>
      </a-breadcrumb>
    </a-space>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'

let $route = useRoute()
const isTopLevelComponent = (path: string) => {
  return $route.matched[0].path === path
}
</script>

<style scoped lang="less"></style>
