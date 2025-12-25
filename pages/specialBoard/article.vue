<template>
  <div class="publish-article-container p-6 bg-white rounded-lg shadow-md">
    <el-page-header content="发布农业知识文章" @back="goBack" />
    
    <el-form
      ref="articleFormRef"
      :model="articleForm"
      :rules="articleRules"
      label-width="100px"
      class="mt-6"
    >
      <!-- 文章标题 -->
      <el-form-item label="文章标题" prop="title">
        <el-input
          v-model="articleForm.title"
          placeholder="请输入文章标题（不超过50字）"
          maxlength="50"
          show-word-limit
          size="large"
        />
      </el-form-item>

      <!-- 文章标签（替换为普通输入框，兼容Element Plus版本） -->
      <el-form-item label="文章标签" prop="tags">
        <el-input
          v-model="tagInputValue"
          placeholder="请输入标签，多个标签用英文逗号分隔"
          size="large"
          @change="handleTagsChange"
        />
        <div class="text-gray-500 text-sm mt-2">最多添加5个标签，用于文章分类检索</div>
      </el-form-item>

      <!-- 文章封面（非必填，补充字段） -->
      <el-form-item label="文章封面" prop="coverImage">
        <el-upload
          class="avatar-uploader"
          action="/api/expert/upload/image"
          :show-file-list="false"
          :on-success="handleCoverSuccess"
          :before-upload="beforeCoverUpload"
        >
          <img v-if="articleForm.coverImage" :src="articleForm.coverImage" class="avatar" />
          <el-icon v-else class="avatar-uploader-icon"><plus /></el-icon>
        </el-upload>
        <div class="text-gray-500 text-sm mt-2">建议尺寸：800*450px，支持jpg/png格式，大小不超过2MB（非必填）</div>
      </el-form-item>

      <!-- 文章摘要（对应 excerpt 字段） -->
      <el-form-item label="文章摘要" prop="excerpt">
        <el-input
          v-model="articleForm.excerpt"
          type="textarea"
          placeholder="请输入文章摘要（不超过200字）"
          maxlength="200"
          show-word-limit
          :rows="3"
          size="large"
        />
      </el-form-item>

      <!-- 文章内容（富文本编辑器） -->
      <el-form-item label="文章内容" prop="content">
        <div class="editor-container border rounded-md">
          <!-- 临时富文本编辑器占位（替换为真实组件） -->
          <textarea
            v-model="articleForm.content"
            class="w-full h-[400px] p-2 border-0 outline-none"
            placeholder="请输入文章内容"
          ></textarea>
        </div>
      </el-form-item>

      <!-- 发布状态 -->
      <el-form-item label="发布状态" prop="status">
        <el-radio-group v-model="articleForm.status" size="large">
          <el-radio label="draft">保存草稿</el-radio>
          <el-radio label="publish">立即发布</el-radio>
        </el-radio-group>
      </el-form-item>

      <!-- 操作按钮 -->
      <el-form-item class="mt-8">
        <el-button type="primary" size="large" @click="submitArticle" :loading="submitLoading">
          {{ articleForm.status === 'draft' ? '保存草稿' : '发布文章' }}
        </el-button>
        <el-button size="large" @click="resetForm" class="ml-4">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, type Ref } from 'vue'
import { 
  ElMessage, ElMessageBox, ElForm, ElFormItem, ElInput, 
  ElRadioGroup, ElRadio, ElButton, ElUpload, ElIcon,
  type FormItemRule // 直接从element-plus核心模块导入
} from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { useRouter } from 'nuxt/app'
import type { PostArticleRequest } from '@/types/knowledgeArticle'

// 临时Mock Pinia Store（解决找不到模块问题，后续替换为真实Store）
const useArticleStore = () => {
  return {
    expertInfo: {
      id: localStorage.getItem('expertId') || 'default-expert-id'
    },
    saveDraft: async (data: PostArticleRequest) => {
      return { code: 200, msg: '草稿保存成功' }
    },
    publishArticle: async (data: PostArticleRequest) => {
      return { code: 200, msg: '文章发布成功' }
    }
  }
}

// 路由实例
const router = useRouter()
// 文章Store
const articleStore = useArticleStore()
// 表单Ref
const articleFormRef = ref<InstanceType<typeof ElForm>>()
// 提交加载状态
const submitLoading = ref(false)
// 标签输入框绑定值
const tagInputValue = ref('')

// 文章表单数据
const articleForm = reactive<Omit<PostArticleRequest, 'author' | 'date'> & {
  coverImage: string
  status: 'draft' | 'publish'
}>({
  title: '',
  excerpt: '',
  content: '',
  tags: [],
  coverImage: '',
  status: 'publish'
})

// 表单校验规则
const articleRules = reactive<Record<string, FormItemRule[]>>({
  title: [
    { required: true, message: '请输入文章标题', trigger: 'blur' },
    { max: 50, message: '标题长度不超过50字', trigger: 'blur' }
  ],
  excerpt: [
    { required: true, message: '请输入文章摘要', trigger: 'blur' },
    { max: 200, message: '摘要长度不超过200字', trigger: 'blur' }
  ],
  content: [
    { required: true, message: '请输入文章内容', trigger: 'blur' }
  ],
  tags: [
    { type: 'array', max: 5, message: '最多添加5个标签', trigger: 'change' }
  ],
  coverImage: []
})

// 标签转换处理
const handleTagsChange = (val: string) => {
  if (!val) {
    articleForm.tags = []
    tagInputValue.value = ''
    return
  }
  const tags = val.split(',').map(tag => tag.trim()).filter(tag => tag)
  articleForm.tags = tags.slice(0, 5)
  tagInputValue.value = articleForm.tags.join(',')
}

// 返回上一页
const goBack = () => {
  router.back()
}

// 封面图片上传前校验
const beforeCoverUpload = (file: File) => {
  const isImage = file.type === 'image/jpeg' || file.type === 'image/png'
  const isLt2M = file.size / 1024 / 1024 < 2

  if (!isImage) {
    ElMessage.error('封面图片只能是 JPG/PNG 格式!')
    return false
  }
  if (!isLt2M) {
    ElMessage.error('封面图片大小不能超过 2MB!')
    return false
  }
  return true
}

// 封面图片上传成功回调
const handleCoverSuccess = (response: any) => {
  if (response.code === 200) {
    articleForm.coverImage = response.data.url
    ElMessage.success('封面上传成功')
  } else {
    ElMessage.error('封面上传失败：' + response.msg)
  }
}

// 重置表单
const resetForm = () => {
  if (articleFormRef.value) {
    articleFormRef.value.resetFields()
    articleForm.coverImage = ''
    articleForm.tags = []
    tagInputValue.value = ''
  }
}

// 提交文章
const submitArticle = async () => {
  try {
    if (!articleFormRef.value) return
    await articleFormRef.value.validate()
    submitLoading.value = true

    // 非空判断修复
    const tags = articleForm.tags && articleForm.tags.length > 0 ? articleForm.tags : undefined

    const submitData: PostArticleRequest = {
      title: articleForm.title,
      excerpt: articleForm.excerpt,
      content: articleForm.content,
      tags: tags,
      author: articleStore.expertInfo.id,
      date: new Date().toISOString()
    }

    const res = articleForm.status === 'draft' 
      ? await articleStore.saveDraft(submitData) 
      : await articleStore.publishArticle(submitData)

    if (res.code === 200) {
      ElMessage.success(articleForm.status === 'draft' ? '草稿保存成功' : '文章发布成功')
      router.push('/expert')
    } else {
      ElMessage.error(res.msg || '操作失败')
    }
  } catch (error) {
    console.error('提交文章失败：', error)
    ElMessage.error('提交失败，请稍后重试')
  } finally {
    submitLoading.value = false
  }
}

// 页面挂载时校验登录状态
onMounted(() => {
  if (!articleStore.expertInfo?.id) {
    ElMessageBox.alert('请先登录专家账号', '权限不足', {
      confirmButtonText: '前往登录',
      type: 'warning'
    }).then(() => {
      router.push('/login')
    })
  }
})
</script>

<style scoped>
.publish-article-container {
  max-width: 1200px;
  margin: 0 auto;
}

.avatar-uploader {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
  width: 200px;
  height: 100px;
}

.avatar-uploader:hover {
  border-color: var(--el-color-primary);
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 100%;
  height: 100%;
  text-align: center;
  line-height: 100px;
}

.avatar {
  width: 100%;
  height: 100%;
  display: block;
  object-fit: cover;
}

.editor-container {
  min-height: 400px;
}
</style>