<template>
    <div class="layout-breadcrumb">
        <a-space direction="vertical">
            <a-breadcrumb :style="{ margin: '16px 0' }" separator=">">
                <a-breadcrumb-item v-for="item in $route.matched" :key="item.path" @click="handleClick(item)" :to="{path: '/'}">
                    <component style="margin: 0 5px;" :is="item.meta.icon"></component>
                    <span>{{ item.meta.title }}</span>
                </a-breadcrumb-item>
            </a-breadcrumb>
        </a-space>
    </div>
</template>


<script setup lang="ts">
import router from '@/router';
import { useRoute } from 'vue-router';

let $route = useRoute();

const handleClick = (item: any) => {
    router.push(item.path)
};
</script>


<style scoped lang="less">
</style>