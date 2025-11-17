<template>
  <div>
    <div class="mb-4">
      <h2 class="text-xl font-semibold">安全设置</h2>
      <p class="text-sm text-gray-500">修改登录密码</p>
    </div>

    <div v-if="notice" :class="['p-3 mb-4 rounded', notice.type === 'success' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">{{ notice.text }}</div>

    <div>
      <label class="block mb-2">邮箱地址</label>
      <div class="flex justify-between w-full">
                <input v-model="security.email" class="w-full border rounded p-2" placeholder="请输入验证码"/>
                <el-button type="success" size="large" :loading="LodingStore.sendLoading" @click="onSendCode(security.email),Countdown(60)" plain>发送验证码</el-button>
      </div>
      <div v-if="LodingStore.sendLoading" class="flex items-center">请在{{ count }}秒后重试</div>
      <label class="block mt-3 mb-2">验证码</label>
      <input v-model="security.code" class="w-full border rounded p-2" placeholder="请输入邮箱验证码" />
      <label class="block mb-2">当前密码</label>
      <input type="password" v-model="security.current" class="w-full border rounded p-2" placeholder="当前密码" />

      <label class="block mt-3 mb-2">新密码</label>
      <input type="password" v-model="security.new" class="w-full border rounded p-2" placeholder="新密码" />

      <label class="block mt-3 mb-2">确认密码</label>
      <input type="password" v-model="security.confirm" class="w-full border rounded p-2" placeholder="确认新密码" />

      <div class="mt-4">
        <button class="btn-primary px-4 py-2 rounded" @click="changePassword">修改密码</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const security = ref({ email:'',code:'',current: '', new: '', confirm: '' })
const notice = ref<{ type: string; text: string } | null>(null)
const LodingStore=useLoadingStore()
function validateEmail(email: string) {
  const re = /^(([^<>()[\\]\\.,;:\s@\\\"]+(\\.[^<>()[\\]\\.,;:\s@\\\"]+)*)|(\\\".+\\\"))@(([^<>()[\\]\\.,;:\s@\\\"]+\\.)+[^<>()[\\]\\.,;:\s@\\\"]{2,})$/i
  return re.test(email)
}
const count = ref(0)
function Countdown(seconds: number) {
    count.value = seconds
    const interval = setInterval(() => {
        if (count.value > 0) {
            count.value--
        } else {
            clearInterval(interval)
            LodingStore.cancelLoading(false)
        }
    }, 1000)
}
async function changePassword() {
  if (security.value.new !== security.value.confirm) {
    notice.value = { type: 'warning', text: '两次输入的密码不一致' }
    return
  }
  if (!security.value.email || !validateEmail(security.value.email)) {
    notice.value = { type: 'warning', text: '请输入正确的邮箱' }
    return
  }
  const isCodeValid = await onVerifyCode(security.value.email, security.value.code)
  if(isCodeValid!==200){
    notice.value = { type: 'warning', text: '验证码错误或已过期' }
    return
  }
  // TODO: 调用 API 修改密码
  notice.value = { type: 'success', text: '密码修改成功（模拟）' }
  security.value = {email:'',code:'',current: '', new: '', confirm: '' }
}
</script>

<style scoped>
.btn-primary { background-color: #0ea5a0; color: white }
.w-full { width: 100% }
.ml-3 { margin-left: .75rem }
</style>
