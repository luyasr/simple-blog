<script setup lang="ts">
import { reactive, ref } from 'vue'
import type { LoginRequest } from '@/types/user'
import { useTokenStore } from '@/stores/modules/token'
import router from '@/router'

const form = reactive<LoginRequest>({ username: '', password: '' })
const loading = ref(false)

let tokenStore = useTokenStore()
const submit = async () => {
  // 登录加载中状态
  loading.value = true
  try {
    await tokenStore.login(form)
    router.push({ name: 'home' })
  } catch (error) {
    console.log(error) 
  } finally {
    loading.value = false
  }
}

const rules = {
  username: [{ required: true, message: '账号是必填项' }],
  password: [{ required: true, message: '密码是必填项' }]
}
</script>

<template>
  <div class="login-wrapper">
    <div class="login-container">
      <a-form :model="form" auto-label-width @submit-success="submit" :rules="rules">
        <a-form-item field="username" :validate-trigger="['change', 'blur']" :hide-asterisk="true">
          <a-input v-model="form.username" icon-user placeholder="请输入用户名" size="large" @press-enter="submit">
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item field="password" :validate-trigger="['change', 'blur']" :hide-asterisk="true">
          <a-input-password v-model="form.password" placeholder="请输入密码" size="large" @press-enter="submit">
            <template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item>
          <a-button html-type="submit" type="primary" style="width: 100%" :loading="loading">登录</a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<style scoped lang="less">
.login-wrapper {
  min-width: 100%;
  min-height: 100vh;
  position: relative;

  .login-container {
    display: flex;
    width: 350px;
    height: 350px;
    position: absolute;
    left: 50%;
    top: 65%;
    transform: translate(-50%, -50%);
  }
}

.a-input,
.a-input-password {
  border-radius: 10px !important;
}
</style>
@/stores/modules/user/index@/types/user@/stores/modules/user/user
