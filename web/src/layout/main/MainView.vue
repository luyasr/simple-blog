<template>
    <div>
        <router-view v-slot="{ Component }">
            <transition name="fade">
                <component :is="Component" v-if="flag" />
            </transition>
        </router-view>
    </div>
</template>

<script setup lang="ts">
import { watch, ref, nextTick } from 'vue'
import { useLayoutSettingStore } from '@/stores/modules/setting'

const flag = ref(true)
let useStore = useLayoutSettingStore()
// 监听仓库内数据是否发生变化, 如果发生变化, 说明用户点击过刷新按钮
watch(() => useStore.refresh, () => {
    // 点击按钮, 路由组件销毁
    flag.value = false
    nextTick(() => {
        // 重新渲染路由组件
        flag.value = true
    })
})
</script>

<style scoped lang="less">
.fade-enter-from {
    opacity: 0;
}

.fade-enter-active {
    transition: all 1s;
}

.fade-enter-to {
    opacity: 1;
}
</style>