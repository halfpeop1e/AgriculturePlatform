<template>
  <div>
    <div class="mb-4">
      <h2 class="text-xl font-semibold">通知设置</h2>
      <p class="text-sm text-gray-500">管理你的通知偏好</p>
    </div>

    <div v-if="notice" :class="['p-3 mb-4 rounded', notice.type === 'success' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">{{ notice.text }}</div>

    <div class="space-y-4">
      <div class="flex items-center gap-3">
        <input type="checkbox" v-model="notifications.email" />
        <span>邮件通知</span>
      </div>
      <div class="flex items-center gap-3">
        <input type="checkbox" v-model="notifications.sms" />
        <span>短信通知</span>
      </div>
      <div class="mt-3">
        <button class="btn-primary px-4 py-2 rounded" @click="saveNotifications">保存通知设置</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const notifications = ref({ email: true, sms: false })
const notice = ref<{ type: string; text: string } | null>(null)

function saveNotifications() {
  // TODO: 调用 API 保存设置
  notice.value = { type: 'success', text: '通知设置已保存' }
}
</script>

<style scoped>
.btn-primary { background-color: #0ea5a0; color: white }
</style>
