<script setup lang="ts">
// 引入左侧菜单子组件
import Logo from './logo/LogoView.vue'
import Menu from './menu/MenuView.vue'
import Tabbar from './tabbar/TabbarView.vue'
import Header from './header/HeaderView.vue'
import Main from './main/MainView.vue'
import { useRouteStore } from '@/stores/modules/route'
// 获取route相关小仓库
let routeStore = useRouteStore()
</script>

<template>
  <a-layout class="layout-container">
    <a-layout-sider collapsible breakpoint="xl">
      <Logo />
      <Menu :menuList="routeStore.menuRoutes"> </Menu>
      <!-- trigger -->
      <template #trigger="{ collapsed }">
        <IconCaretRight v-if="collapsed"></IconCaretRight>
        <IconCaretLeft v-else></IconCaretLeft>
      </template>
    </a-layout-sider>
    <a-layout>
      <a-layout-header>
        <Header />
      </a-layout-header>
      <a-layout style="padding: 0 24px;">
        <Tabbar />
        <Main />
      </a-layout>
    </a-layout>
  </a-layout>
</template>



<style scoped lang="less">
.layout-container {
  width: 100%;
  height: 100vh;
  background: var(--color-fill-1);
  border: 1px solid var(--color-border);
}

.layout-container :deep(.arco-layout-header) {
  height: 64px;
  line-height: 64px;
  background: var(--color-bg-3);
}

.layout-container :deep(.arco-layout-content) {
  color: var(--color-text-2);
  font-weight: 400;
  font-size: 14px;
  background: var(--color-bg-3);
}

.layout-container :deep(.arco-layout-content) {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: var(--color-white);
  font-size: 16px;
  font-stretch: condensed;
  text-align: center;
}
</style>
