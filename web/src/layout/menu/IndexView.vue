<script setup lang='ts'>
// 获取父组件传递过来的路由数组
defineProps(['menuList'])

const onCollapse = (collapsed: boolean) => {
  console.log(collapsed);
};

</script>

<template>
  <div class="menu-demo">
    <a-scrollbar style="height: 100%;overflow: auto;">
      <div style="height: 100%;width: 100vh; background-color: var(--color-primary-light-4);">
        <a-menu :style="{ width: '260px', height: '100%' }" :default-open-keys="['0']" :default-selected-keys="['0_2']"
          show-collapse-button breakpoint="xl" @collapse="onCollapse">
          <template v-for="item in menuList" :key="item.path">
            <a-menu-item v-if="!item.choldren">
              <span>{{ item.path }}</span>
            </a-menu-item>
          </template>
        </a-menu>
      </div>
    </a-scrollbar>
  </div>
</template>

<style scoped lang='less'>
</style>