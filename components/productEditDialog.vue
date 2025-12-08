<template>
  <el-dialog v-model="visibleLocal" modal :closable="true" header="编辑商品信息" :style="{ width: '700px' }">
    <div class="space-y-4">
      <!-- 商品名称 -->
      <div>
        <label class="block text-sm font-medium mb-1">商品名称</label>
        <InputText v-model="form.name" placeholder="请输入商品名称" class="w-full" />
      </div>

      <!-- 分类 -->
      <div>
        <label class="block text-sm font-medium mb-1">分类</label>
        <InputText v-model="form.category" placeholder="请输入分类" class="w-full" />
      </div>

      <!-- 价格与库存 -->
      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium mb-1">价格（¥）</label>
          <InputNumber v-model="form.price" :min="0" :step="0.01" class="w-full" />
        </div>
        <div>
          <label class="block text-sm font-medium mb-1">库存</label>
          <InputNumber v-model="form.stock" :min="0" class="w-full" />
        </div>
      </div>

      <!-- 联系邮箱 -->
      <div>
        <label class="block text-sm font-medium mb-1">联系邮箱</label>
        <InputText v-model="form.contactEmail" type="email" placeholder="请输入联系邮箱" class="w-full" />
      </div>

      <!-- 描述 -->
      <div>
        <label class="block text-sm font-medium mb-1">商品描述</label>
        <Textarea v-model="form.description" rows="4" placeholder="请输入商品描述" class="w-full" />
      </div>

      <!-- 商品图片（展示现有图片，可选择替换）-->
      <div>
        <label class="block text-sm font-medium mb-1">商品图片</label>
        <div v-if="form.images && form.images.length > 0" class="flex gap-2 flex-wrap mb-2">
          <div v-for="(img, idx) in form.images" :key="idx" class="relative w-20 h-20 border rounded">
            <img :src="img" class="w-full h-full object-cover rounded" />
            <Button
              icon="pi pi-times"
              class="p-button-rounded p-button-danger p-button-sm absolute -top-1 -right-1"
              @click="removeImage(idx)"
            />
          </div>
        </div>
        <FileUpload
          mode="basic"
          accept="image/*"
          :multiple="true"
          :auto="false"
          chooseLabel="选择图片"
          class="w-full"
          @select="onImageSelect"
        />
      </div>
    </div>

    <template #footer>
      <div class="flex justify-end gap-2">
        <Button label="取消" severity="secondary" @click="onCancel" />
        <Button label="保存" @click="onSave" />
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import type { productResponse } from '~/types/product'

const props = defineProps<{
  modelValue: boolean
  product: productResponse | null
}>()

const emits = defineEmits<{
  (e: 'update:modelValue', v: boolean): void
  (e: 'save', data: any): void
  (e: 'cancel'): void
}>()

const visibleLocal = ref(props.modelValue)
watch(() => props.modelValue, (v) => (visibleLocal.value = v))
watch(visibleLocal, (v) => emits('update:modelValue', v))

// 表单数据
const form = ref({
  id: '',
  name: '',
  category: '',
  price: 0,
  stock: 0,
  contactEmail: '',
  description: '',
  images: [] as string[],
})

// 当传入 product 时初始化表单
watch(
  () => props.product,
  (p) => {
    if (p) {
      form.value = {
        id: p.id || '',
        name: p.name || '',
        category: p.category || '',
        price: p.price || 0,
        stock: p.stock || 0,
        contactEmail: '', // 如果 product 类型里没有 contactEmail，可从其他字段映射或默认空
        description: p.description || '',
        images: p.image ? p.image.map((img) => img.url) : [],
      }
    }
  },
  { immediate: true }
)

// 移除图片
function removeImage(idx: number) {
  form.value.images.splice(idx, 1)
}

// 选择新图片
function onImageSelect(event: any) {
  const files = event.files as File[]
  if (!files || files.length === 0) return

  files.forEach((file) => {
    const reader = new FileReader()
    reader.onload = (e) => {
      if (e.target?.result) {
        form.value.images.push(e.target.result as string)
      }
    }
    reader.readAsDataURL(file)
  })
}

function onCancel() {
  visibleLocal.value = false
  emits('cancel')
}

function onSave() {
  // 基础校验
  if (!form.value.name.trim()) {
    ElMessage.warning('请输入商品名称')
    return
  }
  if (form.value.price <= 0) {
    ElMessage.warning('价格必须大于0')
    return
  }

  emits('save', { ...form.value })
  visibleLocal.value = false
}
</script>

<style scoped>
/* PrimeVue 组件样式由全局主题控制，这里仅做局部补充 */
</style>
