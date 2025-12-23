<template>
  <!--
    知识卡片组件（简洁版本）
    - 显示文章标签、标题、摘要、作者与日期
    - 使用响应式布局：在小屏为竖向（列），在中等及以上屏幕为横向（行）
    - 样式使用 tailwind 工具类，包含 hover 动画与圆角、阴影
  -->
  <div
    class="bg-white rounded-lg border border-gray-100 shadow-sm hover:shadow-lg transform hover:-translate-y-0.5 transition-all duration-200 overflow-hidden cursor-pointer"
  >
    <div class="flex flex-col md:flex-row">
      <!-- 文本内容区域：包含标签、标题、摘要、作者/日期 -->
      <div class="p-4 flex-1">
        <!-- 标签区域：如果 article.tags 不存在，则使用空数组避免报错 -->
        <div class="flex items-center gap-2 mb-3 flex-wrap">
          <span
            v-for="tag in article.tags || []"
            :key="tag"
            class="inline-flex items-center text-xs px-4 py-0.75 rounded-lg border border-blue-600 text-black bg-white hover:bg-green-50 hover:border-green-200"
          >
            {{ tag }}
          </span>
        </div>

        <!-- 标题：限制两行显示，多余内容用省略号隐藏（通过 .line-clamp-2 实现） -->
        <h3 class="text-lg font-semibold text-gray-800 mb-2 line-clamp-2">{{ article.title }}</h3>

        <!-- 摘要：最多显示三行 -->
        <p class="text-sm text-gray-500 mb-4 line-clamp-3">{{ article.excerpt }}</p>

        <!-- 作者与日期信息 -->
        <div class="flex items-center justify-between text-sm text-gray-400">
          <!-- 使用 truncate 保证长用户名不会撑破布局 -->
          <div class="truncate">{{ article.author }}</div>
          <div>{{ article.date }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// 说明：组件仅通过 props 接收一个 article 对象并渲染。
// 如果后续需要交互（点击跳转、收藏等），可在这里添加 emits 或 useRouter 跳转逻辑。
import type { KnowledgeArticle } from '~/types/knowledgeArticle'

// 接收外部传入的 article，类型由项目全局类型定义（types/knowledgeArticle）提供
const props = defineProps<{ article: KnowledgeArticle }>()

// 小提示：如果想对 title 或 excerpt 做防 XSS 的额外处理，可以在这里使用 DOMPurify
// 或在后端返回已安全化的文本。当前组件假设 article 字段为安全的纯文本。
</script>

<style scoped>
/*
  辅助的行截断样式（line-clamp）。
  Tailwind 的官方插件也提供 line-clamp 能力，这里是基于 -webkit-box 的手工实现，兼容大部分现代浏览器。
*/
.line-clamp-2 {
  display: -webkit-box;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
.line-clamp-3 {
  display: -webkit-box;
  line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
