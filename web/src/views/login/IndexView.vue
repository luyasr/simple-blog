<script setup lang='ts'>
import { reactive, ref } from 'vue'
import { Message } from '@arco-design/web-vue'
import type { LoginForm } from '@/types/user'
import { useUserStore } from '@/stores/modules/user/index'
import router from '@/router';

const form = reactive<LoginForm>({ username: "", password: "" })
const loading = ref(false)

let userStore = useUserStore()
const submit = async () => {
  // 登录加载中状态
  loading.value = true
  try {
    await userStore.login(form)
    router.push('/')
  } catch (error) {
    Message.error(`${error}`)
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
          <a-input v-model="form.username" icon-user placeholder="请输入用户名" @press-enter="submit" allow-clear>
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item field="password" :validate-trigger="['change', 'blur']" :hide-asterisk="true">
          <a-input-password v-model="form.password" placeholder="请输入密码" @press-enter="submit" allow-clear>
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

<style scoped lang='less'>
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
</style>@/stores/modules/user/index@/types/user@/stores/modules/user/user