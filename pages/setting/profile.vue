<template>
  <div>
    <div class="mb-4">
      <h2 class="text-xl font-semibold">个人信息</h2>
      <p class="text-sm text-gray-500">编辑你的昵称、邮箱和个人简介</p>
    </div>

    <div v-if="notice" :class="['p-3 mb-4 rounded', notice.type === 'success' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800']">{{ notice.text }}</div>

    <div class="p-fluid">
      <label class="block mb-2">昵称</label>
      <input v-model="profile.nickname" class="w-full border rounded p-2" placeholder="请输入昵称" />

      <label class="block mt-3 mb-2">个人头像 URL</label>
      <input v-model="profile.avatar" class="w-full border rounded p-2" placeholder="头像图片地址 (可选)" />

      <label class="block mt-3 mb-2">个人简介</label>
      <textarea v-model="profile.bio" rows="4" class="w-full border rounded p-2"></textarea>

      <label class="block mt-3 mb-2">所在地区</label>
      <input v-model="profile.location" class="w-full border rounded p-2" placeholder="例如：北京市·通州区" />

      <label class="block mt-3 mb-2">联系方式（手机号）</label>
      <input v-model="profile.phone" class="w-full border rounded p-2" placeholder="例如：+86 138-0000-0000" />

      <label class="block mt-3 mb-2">联系地址</label>
      <input v-model="profile.address" class="w-full border rounded p-2" placeholder="详细地址（可选）" />

      <label class="block mt-3 mb-2">标签（逗号分隔）</label>
      <input v-model="tagsInput" class="w-full border rounded p-2" placeholder="例如：前端工程师,农业物联网" />

      <div class="mt-4">
        <button class="btn-primary px-4 py-2 rounded" @click="saveProfile">保存</button>
        <button class="ml-3 px-4 py-2 rounded" @click="resetProfile">重置</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

const profile = ref({ nickname: '', email: '', bio: '', avatar: '', location: '', phone: '', address: '', tags: [] as string[] })
const tagsInput = ref('')

onMounted(() => {
  if (Array.isArray(profile.value.tags) && profile.value.tags.length) {
    tagsInput.value = profile.value.tags.join(',')
  }
})
const notice = ref<{ type: string; text: string } | null>(null)

function validateEmail(email: string) {
  const re = /^(([^<>()[\\]\\.,;:\s@\\\"]+(\\.[^<>()[\\]\\.,;:\s@\\\"]+)*)|(\\\".+\\\"))@(([^<>()[\\]\\.,;:\s@\\\"]+\\.)+[^<>()[\\]\\.,;:\s@\\\"]{2,})$/i
  return re.test(email)
}

function saveProfile() {
  if (!profile.value.nickname) {
    notice.value = { type: 'warning', text: '昵称不能为空' }
    return
  }
  if (!profile.value.email || !validateEmail(profile.value.email)) {
    notice.value = { type: 'warning', text: '请输入正确的邮箱' }
    return
  }
  // 将 tagsInput 转换为数组
  profile.value.tags = tagsInput.value
    .split(',')
    .map((t) => t.trim())
    .filter((t) => t.length > 0)
  // TODO: 调用 API 保存
  notice.value = { type: 'success', text: '个人信息已保存' }
}

function resetProfile() {
  profile.value = { nickname: '', email: '', bio: '', avatar: '', location: '', phone: '', address: '', tags: [] }
  tagsInput.value = ''
  notice.value = null
}
</script>

<style scoped>
.btn-primary { background-color: #0ea5a0; color: white }
.ml-3 { margin-left: .75rem }
.w-full { width: 100% }
</style>
