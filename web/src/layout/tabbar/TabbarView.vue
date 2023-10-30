<template>
  <div class="layout-breadcrumb">
    <a-space direction="vertical">
      <a-breadcrumb :style="{ margin: '16px 0' }">
        <a-breadcrumb-item v-for="item in $route.matched" :key="item.path">
          <template #separator>
            <span v-if="item.meta.title">></span>
          </template>
          <span>
            <router-link v-if="item.meta.title" :to="item.path">
              {{ item.meta.title }}
            </router-link>
          </span>
        </a-breadcrumb-item>
      </a-breadcrumb>
    </a-space>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'

let $route = useRoute()
</script>

<style scoped lang="less"></style>
