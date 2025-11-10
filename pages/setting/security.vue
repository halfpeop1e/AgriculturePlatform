<template>
  <div>
    <div class="mb-4">
      <h2 class="text-xl font-semibold">安全设置</h2>
      <p class="text-sm text-gray-500">修改登录密码</p>
    </div>

    <div v-if="notice" :class="['p-3 mb-4 rounded', notice.type === 'success' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">{{ notice.text }}</div>

    <div>
       <label class="block mt-3 mb-2">邮箱</label>
      <input v-model="security.email" class="w-full border rounded p-2" placeholder="请输入邮箱" />

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
const security = ref({ email:'',current: '', new: '', confirm: '' })
const notice = ref<{ type: string; text: string } | null>(null)

function changePassword() {
  if (security.value.new !== security.value.confirm) {
    notice.value = { type: 'warning', text: '两次输入的密码不一致' }
    return
  }
  // TODO: 调用 API 修改密码
  notice.value = { type: 'success', text: '密码修改成功（模拟）' }
  security.value = {email:'',current: '', new: '', confirm: '' }
}
</script>

<style scoped>
.btn-primary { background-color: #0ea5a0; color: white }
.w-full { width: 100% }
.ml-3 { margin-left: .75rem }
</style>
