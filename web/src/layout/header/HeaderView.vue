<template>
    <div class="layout-header">
        <a-button class="layut-header-button" type="dashed" size="large" shape="round" @click="handleRefresh">
            <template #icon>
                <icon-refresh />
            </template>
        </a-button>
        <a-button class="layut-header-button" type="dashed" size="large" shape="round" @click="toggle">
            <template #icon>
                <icon-fullscreen />
            </template>
        </a-button>
        <a-button class="layut-header-button" type="dashed" size="large" shape="round">
            <template #icon>
                <icon-settings />
            </template>
        </a-button>
        <a-dropdown trigger="hover" show-arrow>
            <a-space>
                <a-avatar :size="35" shape="square">
                    <img alt="avatar" :src="userStore.avatar" />
                </a-avatar>
                <icon-caret-down />
            </a-space>
            <template #content>
                <a-doption>个人中心</a-doption>
                <a-doption @click="handleLogout">退出登录</a-doption>
            </template>
        </a-dropdown>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Message } from '@arco-design/web-vue'
import { useFullscreen } from '@vueuse/core'
import { useUserStore } from '@/stores/modules/user'
import { useLayoutSettingStore } from '@/stores/modules/setting'
import router from '@/router'

let userStore = useUserStore()
let layoutSettingStore = useLayoutSettingStore()
const handleRefresh = () => {
    layoutSettingStore.refresh = !layoutSettingStore.refresh
}
// 反应式全屏 https://vueuse.org/core/useFullscreen/
const el = ref()
const { toggle } = useFullscreen(el)

const handleLogout = async () => {
    try {
        await userStore.logout()
        router.push({ name: 'login' })
    } catch (error) {
        Message.error(`${error}`)
    }
}
</script>


<style scoped lang="less">
.layout-header {
    height: 100%;
    display: flex;
    justify-content: flex-end;
    align-items: center;
}

.layut-header-button {
    background: var(--color-fill-1);
}

.layout-header>* {
    /* 设置内容间隔为 10 像素 */
    margin-right: 10px;
}

.hover {
    color: red;
}
</style>