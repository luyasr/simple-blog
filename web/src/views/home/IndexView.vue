<script setup lang='ts'>
import { onMounted } from 'vue';
import { useUserStore } from '@/stores/modules/user' 

onMounted(() => {
  let userStore = useUserStore();
  userStore.userInfo()
});
</script>

<template>

  <div>
    <h1>首页</h1>
  </div>

</template>

<style scoped lang='less'>

</style>