<script setup lang='ts'>
import settings from '@/config/settings'
</script>

<template>
  <div class="logo" v-if="!settings.hideInLogo">
    <img :src="settings.logo" alt="">
    <p>{{ settings.title }}</p>
  </div>
</template>

<style scoped lang='less'>
.logo {
  display: flex;
  width: 100%;
  height: @base-menu-logo-height;
  color: white;
  align-items: center;
  padding: 20px;

  img {
    width: 40px;
    height: 40px;
  }

  p {
    font-size: @base-menu-title-font-size;
    margin-left: 10px;
  }
}
</style>