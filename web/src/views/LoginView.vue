<script setup lang='ts'>
import { reactive, ref } from 'vue'
import { Message } from '@arco-design/web-vue'
import type { LoginForm } from '@/types/token'
import { useTokenStore } from '@/stores/modules/token'
import icon from '@arco-design/web-vue/es/icon';

const form = reactive<LoginForm>({ username: "", password: "" })
const loading = ref(false)

let tokenStore = useTokenStore()
const submit = async () => {
  // 登录加载中状态
  loading.value = true
  try {
    await tokenStore.login(form)
    Message.info('登录成功')
    loading.value = false
  } catch (error) {
    loading.value = false
    Message.error(`${error}`)
  }
}
</script>

<template>
  <div class="login-wrapper">
    <div class="login-container">
      <a-form :model="form" auto-label-width @submit="submit">
        <div class="field">
          <a-form-item field="username">
            <a-input v-model="form.username" icon-user placeholder="请输入用户名" @press-enter="submit" allow-clear>
              <template #prefix>
                <icon-user />
              </template>
            </a-input>
          </a-form-item>
        </div>
        <div class="field">
          <a-form-item field="password">
            <a-input-password v-model="form.password" placeholder="请输入密码" @press-enter="submit" allow-clear>
              <template #prefix>
                <icon-lock />
              </template>
            </a-input-password>
          </a-form-item>
        </div>
        <div class="field">
          <a-form-item>
            <a-button html-type="submit" type="primary" style="width: 100%" :loading="loading">登录</a-button>
          </a-form-item>
        </div>
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
    border-radius: 20px;
    width: 350px;
    height: 350px;
    position: absolute;
    left: 50%;
    top: 65%;
    transform: translate(-50%, -50%);
    justify-content: flex-start;

    .field {
      justify-content: flex-start;
    }
  }
}
</style>