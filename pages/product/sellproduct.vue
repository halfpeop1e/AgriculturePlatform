<template>
        <section class="py-8 px-4">
            <div class="max-w-3xl mx-auto">
                <div class="bg-white/80 dark:bg-slate-900/80 rounded-2xl shadow-lg p-6">
                    <h1 class="text-2xl font-semibold mb-4">上架商品</h1>

                    <Toast />

                    <form @submit.prevent="onSubmit" class="space-y-4">
                        <!-- 商品标题 -->
                        <div>
                            <label class="block text-sm font-medium mb-1">商品标题</label>
                            <InputText v-model="form.name" placeholder="请输入商品标题" class="w-full" />
                            <p v-if="errors.name" class="text-red-500 text-sm mt-1">{{ errors.name }}</p>
                        </div>
                        <div>
                            <label class="block text-sm font-medium mb-1">商品描述</label>
                            <InputText v-model="form.description" placeholder="请输入商品描述" class="w-full" />
                            <p v-if="errors.description" class="text-red-500 text-sm mt-1">{{ errors.description }}</p>
                        </div>
                        <!-- 分类 -->
                        <div>
                            <label class="block text-sm font-medium mb-1">分类</label>
                            <Dropdown v-model="form.category" :options="categories" optionLabel="label" optionValue="value" placeholder="选择分类" class="w-full" />
                            <p v-if="errors.category" class="text-red-500 text-sm mt-1">{{ errors.category }}</p>
                        </div>

                        <!-- 价格 -->
                        <div>
                            <label class="block text-sm font-medium mb-1">价格 (¥)</label>
                            <InputNumber v-model="form.price" mode="decimal" :min="0" class="w-full" />
                            <p v-if="errors.price" class="text-red-500 text-sm mt-1">{{ errors.price }}</p>
                        </div>

                        <!-- 数量 -->
                        <div>
                            <label class="block text-sm font-medium mb-1">数量</label>
                            <InputNumber v-model="form.stock" :min="1" class="w-full" />
                            <p v-if="errors.stock" class="text-red-500 text-sm mt-1">{{ errors.stock }}</p>
                        </div>
                        <!-- 联系邮箱 -->
                        <div>
                            <label class="block text-sm font-medium mb-1">联系方式（邮箱）</label>
                            <InputText v-model="form.contactEmail" placeholder="用于买家联系" class="w-full" />
                            <p v-if="errors.contactEmail" class="text-red-500 text-sm mt-1">{{ errors.contactEmail }}</p>
                        </div>

                        <!-- 图片上传 -->
                        <div>
                            <label class="block text-sm font-medium mb-1">商品图片</label>
                                        <div class="mb-2">
                                            <input ref="fileInput" type="file" accept="image/*" multiple @change="onFilesChange" class="hidden" />
                                            <Button label="选择图片" icon="pi pi-upload" class="mr-2" @click="triggerFileSelect" />
                                            <Button label="清空图片" icon="pi pi-trash" class="p-button-secondary" @click="clearFiles" />
                                        </div>
                            <div class="grid grid-cols-3 gap-3">
                                <div v-for="(f, idx) in fileList" :key="idx" class="relative border rounded-md overflow-hidden">
                                    <img :src="f.url" class="w-full h-32 object-cover" />
                                    <button type="button" @click="removeFile(idx)" class="absolute top-1 right-1 bg-white/70 rounded-full p-1 text-sm">✕</button>
                                </div>
                            </div>
                            <p v-if="errors.images" class="text-red-500 text-sm mt-1">{{ errors.images }}</p>
                            <p class="text-xs text-slate-400 mt-2">建议尺寸：800x800，单张不超过 2MB</p>
                        </div>

                        <div class="flex gap-3 mt-4">
                            <Button type="submit" label="发布商品" class="p-button-primary" />
                            <Button label="重置" class="p-button-secondary" @click="onReset" />
                        </div>
                    </form>
                </div>
            </div>
        </section>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useToast } from 'primevue/usetoast'
import type { postProductRequest } from '@/types/product'
definePageMeta({ layout: 'home-page-layout' })

const toast = useToast()

const fileInput = ref<HTMLInputElement | null>(null)

const form = reactive<postProductRequest>({
    name: '',
    category: '' as string,
    price: 0 as number,
    stock: null as any,
    description: '',
    images: [] as string[],
    saler:'' as string ,
    contactEmail: '',
    salerId:''
})

const categories = [
    { label: '粮食', value: 'grain' },
    { label: '蔬菜', value: 'vegetable' },
    { label: '水果', value: 'fruit' },
    { label: '种苗', value: 'seedling' }
]

const fileList = ref<{ name: string; url: string }[]>([])

const errors = reactive<Record<string, string | null>>({
    name: null,
    category: null,
    price: null,
    stock: null,
    description: null,
    contactEmail: null,
    saler:null,
    images: null
})

function validate() {
    let ok = true
    // reset errors
    Object.keys(errors).forEach((k) => (errors[k] = null))

    if (!form.name || !form.name.trim()) {
        errors.name = '请输入商品标题'
        ok = false
    }
    if (!form.category) {
        errors.category = '请选择分类'
        ok = false
    }
    if (form.price == null || Number(form.price) < 0) {
            errors.price = '请输入有效价格'
        ok = false
    }
    if (!form.stock || Number(form.stock) < 1) {
        errors.stock = '请输入有效数量'
        ok = false
    }
    if (!form.description || !form.description.trim()) {
        errors.description = '请填写商品描述'
        ok = false
    }
    if (!form.saler || !form.saler.trim()) {
        errors.saler = '请填写商家名称'
        ok = false
    }
    if (!form.contactEmail || !/^[^@\s]+@[^@\s]+\.[^@\s]+$/.test(form.contactEmail)) {
        errors.contactEmail = '请输入有效邮箱'
        ok = false
    }
    if (fileList.value.length === 0) {
        errors.images = '请上传至少一张商品图片'
        ok = false
    }

    return ok
}

function onFilesChange(e: Event) {
    const input = e.target as HTMLInputElement
    if (!input.files) return
    const files = Array.from(input.files)

    files.forEach((file) => {
        if (!file.type.startsWith('image/')) {
            toast.add({ severity: 'error', summary: '文件错误', detail: '仅支持图片文件', life: 3000 })
            return
        }
        if (file.size / 1024 / 1024 >= 10) {
            toast.add({ severity: 'error', summary: '文件过大', detail: '每张图片需小于 10MB', life: 3000 })
            return
        }

        const reader = new FileReader()
        reader.onload = (ev) => {
            const url = ev.target?.result as string
            fileList.value.push({ name: file.name, url })
            form.images = fileList.value.map((f) => f.url)
        }
        reader.readAsDataURL(file)
    })

    // clear native input to allow reselect same file
    if (fileInput.value) fileInput.value.value = ''
}

    function triggerFileSelect() {
        fileInput.value?.click()
    }

function removeFile(idx: number) {
    fileList.value.splice(idx, 1)
    form.images = fileList.value.map((f) => f.url)
}

function clearFiles() {
    fileList.value = []
    form.images = []
}

async function onSubmit() {
    if (!validate()) {
        toast.add({ severity: 'warn', summary: '表单未通过', detail: '请检查必填项', life: 3000 })
        return
    }

    try {
        const userStore = useUserStore()
        toast.add({ severity: 'info', summary: '提交中', detail: '正在上架商品...', life: 2000 })
        form.saler=userStore.userinfo.nickname
        form.salerId=userStore.userId
        console.log('提交的商品信息：', form)
        const response=await PostProduct(form)
        if(response?.status!==200){
            throw new Error('商品上架失败，请稍后重试')
        }
        else{
            toast.add({ severity: 'success', summary: '完成', detail: '商品已上架', life: 3000 })
            onReset()
        }
        
    } catch (err) {
        toast.add({ severity: 'error', summary: '提交失败', detail: String(err), life: 3000 })
    }
}

function onReset() {
    form.name = ''
    form.category = ''
    form.price = 0
    form.stock = 1
    form.description = ''
    form.saler=''
    form.contactEmail = ''
    clearFiles()
    Object.keys(errors).forEach((k) => (errors[k] = null))
}
</script>

<style scoped>
/* 小量样式使上传区域居中 */
.el-upload .el-upload__input { display: none; }
</style>
