<script setup lang='ts'>

</script>

<template>
  <div class="layout-container">
    <!-- 菜单栏 -->
    <div class="layout-menu">菜单栏</div>
    <!-- 导航栏 -->
    <div class="layout-nav">导航栏</div>
    <!-- 内容 -->
    <div class="layout-body">
      内容
      <p style="height: 2000px;;"></p>
    </div>
  </div>
</template>

<style scoped lang='less'>
.layout-container {
  width: 100%;
  height: 100vh;
  background-color: antiquewhite;

  .layout-menu {
    position: fixed;
    width: @base-menu-width;
    height: 100vh;
    background-color: aqua;
  }

  .layout-nav {
    position: fixed;
    width: calc(100% - @base-menu-width);
    height: @base-nav-height;
    top: 0px;
    left: @base-menu-width;
    background-color: aquamarine;
  }

  .layout-body {
    position: absolute;
    width: calc(100% - @base-menu-width);
    height: calc(100vh - @base-nav-height);
    left: @base-menu-width;
    top: @base-nav-height;
    padding: 20px;
    overflow: auto;
  }
}
</style>