<!--
  productEditDialog.vue
  说明：商品编辑弹窗（PrimeVue Dialog + TailwindCSS）
  Props:
    - modelValue: 控制弹窗显示
    - product: 当前商品信息
  Emits:
    - update:modelValue
    - save
    - cancel
-->
<template>
  <Dialog
    v-model:visible="visible"
    header="编辑商品信息"
    modal
    class="w-full max-w-3xl"
    :draggable="false"
  >
    <div class="space-y-6">
      <div class="grid gap-4 md:grid-cols-2">
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-700">商品名称</label>
          <InputText v-model.trim="form.name" placeholder="请输入商品名称" class="w-full" />
        </div>

        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-700">分类</label>
          <InputText v-model.trim="form.category" placeholder="请输入分类" class="w-full" />
        </div>
      </div>

      <div class="grid gap-4 md:grid-cols-2">
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-700">价格（¥）</label>
          <InputNumber v-model="form.price" :min="0" :step="0.01" class="w-full" inputClass="w-full" />
        </div>
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-700">库存</label>
          <InputNumber v-model="form.stock" :min="0" class="w-full" inputClass="w-full" />
        </div>
      </div>

      <div class="space-y-2">
        <label class="text-sm font-medium text-slate-700">联系邮箱</label>
        <InputText v-model.trim="form.contactEmail" type="email" placeholder="请输入联系邮箱" class="w-full" />
      </div>

      <div class="space-y-2">
        <label class="text-sm font-medium text-slate-700">商品描述</label>
        <Textarea v-model.trim="form.description" rows="4" placeholder="请输入商品描述" class="w-full" />
      </div>

      <div class="space-y-3">
        <label class="text-sm font-medium text-slate-700">商品图片</label>
        <div v-if="form.images.length" class="flex flex-wrap gap-3">
          <div
            v-for="(img, idx) in form.images"
            :key="`${img}-${idx}`"
            class="relative h-24 w-24 overflow-hidden rounded-lg border border-slate-200"
          >
            <img :src="img" alt="商品图片" class="h-full w-full object-cover" />
            <Button
              icon="pi pi-times"
              severity="danger"
              size="small"
              rounded
              class="absolute -right-2 -top-2"
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
        <p class="text-xs text-slate-500">支持多张上传，将自动追加至图片列表</p>
      </div>
    </div>

    <template #footer>
      <div class="flex justify-end gap-3">
        <Button label="取消" severity="secondary" @click="onCancel" />
        <Button label="保存" @click="onSave" />
      </div>
    </template>
  </Dialog>
</template>

<script setup lang="ts">
import { ref, watch, computed } from 'vue'
import Dialog from 'primevue/dialog'
import InputText from 'primevue/inputtext'
import InputNumber from 'primevue/inputnumber'
import Textarea from 'primevue/textarea'
import FileUpload from 'primevue/fileupload'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'
import type { productResponse } from '~/types/product'

const props = defineProps<{
  modelValue: boolean
  product: productResponse | null
}>()

type ProductFormPayload = {
  id: string
  name: string
  category: string
  price: number
  stock: number
  contactEmail: string
  description: string
  images: string[]
}

const emits = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'save', value: ProductFormPayload): void
  (e: 'cancel'): void
}>()

const toast = useToast()

const visible = computed({
  get: () => props.modelValue,
  set: (value: boolean) => emits('update:modelValue', value)
})

const form = ref<ProductFormPayload>({
  id: '',
  name: '',
  category: '',
  price: 0,
  stock: 0,
  contactEmail: '',
  description: '',
  images: []
})

watch(
  () => props.product,
  (product) => {
    if (!product) return
    form.value = {
      id: product.id || '',
      name: product.name || '',
      category: product.category || '',
      price: product.price || 0,
      stock: product.stock || 0,
      contactEmail: '',
      description: product.description || '',
      images: product.image ? product.image.map((img) => img.url) : []
    }
  },
  { immediate: true }
)

const getFormPayload = (): ProductFormPayload => ({ ...form.value })

const removeImage = (index: number) => {
  form.value.images.splice(index, 1)
}

const onImageSelect = (event: { files?: File[] }) => {
  const files = event.files
  if (!files?.length) return

  files.forEach((file) => {
    const reader = new FileReader()
    reader.onload = (e) => {
      if (typeof e.target?.result === 'string') {
        form.value.images.push(e.target.result)
      }
    }
    reader.readAsDataURL(file)
  })
}

const onCancel = () => {
  emits('cancel')
  visible.value = false
}

const onSave = () => {
  if (!form.value.name.trim()) {
    toast.add({ severity: 'warn', summary: '提醒', detail: '请输入商品名称', life: 3000 })
    return
  }
  if (form.value.price <= 0) {
    toast.add({ severity: 'warn', summary: '提醒', detail: '价格必须大于 0', life: 3000 })
    return
  }

  emits('save', getFormPayload())
  visible.value = false
}
</script>

<style scoped>
.p-dialog .p-dialog-header {
  border-bottom: none;
}
</style>
