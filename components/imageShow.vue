<!--
  imageShow.vue
  说明：轮播展示组件（基于 Element Plus 的 el-carousel）
  Props:
    - size: 轮播高度，例如 '350px'
    - images: 数组，元素形如 { url: string, text: string }
  行为：在每个轮播项上渲染背景图与覆盖文本，适合首页大图轮播。
-->
<template>
    <el-carousel :interval="4000" :height="props.size">
         <el-carousel-item
      v-for="(item, index) in images"
      :key="index"
      class="relative"
    >
      <!-- 背景图片 -->
      <img
        :src="item.url"
        alt=""
        class="w-full h-full object-cover rounded-lg"
        
      />

      <!-- 居中文字 -->
      <div
        v-if="item.text"
        class="absolute inset-0 flex items-center justify-center bg-black/40 text-white text-5xl font-semibold"
        @click="imgNavigate(item.link as string)"
      >
        {{ item.text }}
      </div>
    </el-carousel-item>
    </el-carousel>
</template>

<script setup lang="ts">
interface ImageItem {
  url: string
  // 文本是可选的：某些轮播项可能只展示图片不需要覆盖文字
  text?: string
  link?: string
}
import type { PropType } from 'vue'

const props = defineProps({
  size: {
    type: String,
    default: '350px',
  },
  images: {
    type: Array as PropType<ImageItem[]>,
    default: () => [] as ImageItem[],
    // 示例结构：
    // [
    //   { url: 'https://example.com/1.jpg', text: '图片一描述' },
    //   { url: 'https://example.com/2.jpg' }, // 没有 text 时也能正常展示
    // ]
  },
})
const imgNavigate = (link: string) => {
  if (link) {
    console.log("ping")
    navigateTo(link)
  }
  else {
    return
  }
}
</script>

<style scoped>
.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>
