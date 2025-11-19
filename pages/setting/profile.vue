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

      <label class="block mt-3 mb-2">个人头像</label>
      <!-- 使用 PrimeVue FileUpload 替换原有输入，限制最多 1 张；选中后打开裁剪对话框 -->
      <div class="mb-2">
        <div v-if="profile.avatar" class="mb-2">
          <img :src="profile.avatar" alt="avatar" class="w-24 h-24 rounded-full object-cover border" />
        </div>
        <!-- PrimeVue FileUpload（需在项目中注册 PrimeVue） -->
        <FileUpload
          mode="basic"
          name="avatar"
          accept="image/*"
          :maxFileSize="10 * 1024 * 1024"
          :auto="false"
          :multiple="false"
          chooseLabel="选择图片"
          @select="onFileSelect"
          class="w-full"
        />
      </div>

      <!-- 裁剪对话框（使用 cropperjs 动态导入） -->
      <Dialog v-model:visible="showCropper" header="裁剪头像" :modal="true" :style="{width: '600px'}">
        <div class="cropper-container" v-if="selectedImageUrl">
          <img ref="cropperImage" :src="selectedImageUrl" alt="to-crop" style="max-width:100%; display:block; margin:0 auto;" />
        </div>
        <template #footer>
          <Button label="取消" class="p-button-text" @click="closeCropper" />
          <Button label="保存" class="p-button-primary" @click="saveCropped" />
        </template>
      </Dialog>

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
import { ref, onMounted, nextTick } from 'vue'
import FileUpload from 'primevue/fileupload'
import Dialog from 'primevue/dialog'
import Button from 'primevue/button'
import 'cropperjs/dist/cropper.css'
// cropperjs will be dynamically imported when needed

const profile = ref({ nickname: '', bio: '', avatar: '', location: '', phone: '', address: '', tags: [] as string[] })
const tagsInput = ref('')
const showCropper = ref(false)
const selectedFile = ref<File | null>(null)
const selectedImageUrl = ref<string | null>(null)
const cropperInstance: { value: any } = ref(null)
const cropperImage = ref<HTMLImageElement | null>(null)

onMounted(() => {
  if (Array.isArray(profile.value.tags) && profile.value.tags.length) {
    tagsInput.value = profile.value.tags.join(',')
  }
})
const notice = ref<{ type: string; text: string } | null>(null)



function saveProfile() {
  if (!profile.value.nickname) {
    notice.value = { type: 'warning', text: '昵称不能为空' }
    return
  }
  
  // 将 tagsInput 转换为数组
  profile.value.tags = tagsInput.value
    .split(',')
    .map((t) => t.trim())
    .filter((t) => t.length > 0)
  // TODO: 调用 API 保存
  changeUserProfile(profile.value)
  console.log('保存的个人信息：', profile.value)
  notice.value = { type: 'success', text: '个人信息已保存' }
}

// PrimeVue FileUpload select 事件处理：读取文件并打开裁剪对话框
function onFileSelect(e: any) {
  const file = e.files && e.files.length ? e.files[0] : null
  if (!file) return
  selectedFile.value = file
  const reader = new FileReader()
  reader.onload = (ev) => {
    selectedImageUrl.value = ev.target?.result as string
    // 打开裁剪对话框，稍后动态加载 cropper 并初始化
    showCropper.value = true
    nextTick(async () => {
      try {
        const Cropper = (await import('cropperjs')).default
        if (cropperInstance.value) {
          cropperInstance.value.destroy()
          cropperInstance.value = null
        }
        if (cropperImage.value) {
          // 将 options 断言为 any 以避免类型声明不匹配时的对象字面量错误
          const cropperOptions: any = {
            aspectRatio: 1, // 裁剪为正方形
            viewMode: 1,
            autoCropArea: 0.8,
            movable: true,
            zoomable: true,
            background: false,
          }
          cropperInstance.value = new Cropper(cropperImage.value as HTMLImageElement, cropperOptions)
        }
      } catch (err) {
        // cropperjs 未安装时降级：直接使用原图
        console.warn('cropperjs 未安装，跳过裁剪：', err)
      }
    })
  }
  reader.readAsDataURL(file)
}

function closeCropper() {
  showCropper.value = false
  selectedFile.value = null
  selectedImageUrl.value = null
  if (cropperInstance.value) {
    cropperInstance.value.destroy()
    cropperInstance.value = null
  }
}

async function saveCropped() {
  try {
    if (cropperInstance.value) {
      const canvas = cropperInstance.value.getCroppedCanvas({ width: 300, height: 300 })
      const dataUrl = canvas.toDataURL('image/png')
      profile.value.avatar = dataUrl
    } else if (selectedImageUrl.value) {
      // 没有 cropper 时直接使用选中的图片
      profile.value.avatar = selectedImageUrl.value
    }
  } catch (err) {
    console.error(err)
  } finally {
    closeCropper()
  }
}

function resetProfile() {
  profile.value = { nickname: '', bio: '', avatar: '', location: '', phone: '', address: '', tags: [] }
  tagsInput.value = ''
  notice.value = null
}
</script>

<style scoped>
.btn-primary { background-color: #0ea5a0; color: white }
.ml-3 { margin-left: .75rem }
.w-full { width: 100% }
</style>
