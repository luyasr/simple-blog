<script setup lang='ts'>
import { reactive } from 'vue'
import { Message } from '@arco-design/web-vue'
import type { LoginForm } from '../api/token/type'
import { useTokenStore } from '../stores/modules/token'

const form = reactive<LoginForm>({ username: "", password: "" })

let tokenStore = useTokenStore()
const submit = async () => {
  try {
    await tokenStore.login(form)
    Message.info('登录成功')
  } catch (error) {
    Message.error(`${error}`)
  }
}
</script>

<template>
  <div class="login-wrapper">
    <div class="login-container">
      <a-form :model="form" @submit="submit">
        <a-form-item field="username">
          <a-input v-model="form.username" icon-user placeholder="请输入用户名" allow-clear>
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item field="password">
          <a-input-password v-model="form.password" placeholder="请输入密码" allow-clear>
            <template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item>
          <a-button html-type="submit" type="primary">登录</a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<style scoped lang='less'>
.login-wrapper {
  min-width: 100%;
  min-height: 100vh;
  position: relative;
}

.login-container {
  text-align: center;
  border-radius: 20px;
  width: 400px;
  height: 400px;
  position: absolute;
  left: 50%;
  top: 65%;
  transform: translate(-50%, -50%);
}
</style>