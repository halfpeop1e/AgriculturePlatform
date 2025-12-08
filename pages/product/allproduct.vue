<template>
  <div>
    <div v-if="initialLoading" class="space-y-6">
      <div class="m-10"></div>
      <el-skeleton :rows="4" animated />
      <div class="m-10"></div>
      <el-skeleton :rows="4" animated />
      <div class="m-10"></div>
      <el-skeleton :rows="4" animated />
      <div class="m-10"></div>
      <el-skeleton :rows="4" animated />
    </div>

    <div v-else class="product-list-container">
      <ul
        v-infinite-scroll="loadMore"
        class="product-list"
        :infinite-scroll-disabled="loading || finished"
        infinite-scroll-distance="160"
      >
        <li
          v-for="(item, index) in products"
          :key="item.id || index"
          class="product-item"
        >
          <ProductCard :product="item" @open-buy-dialog="dialogControler.openDialog(item)" />
        </li>
      </ul>

      <el-empty
        v-if="!products.length && !loading"
        description="暂时没有上架的助农产品"
        class="mt-6"
      />

      <el-skeleton
        v-else-if="loading"
        :rows="3"
        animated
        class="mt-6"
      />

      <div
        v-else-if="finished && products.length"
        class="text-center text-gray-400 mt-4 text-sm"
      >
        已加载全部 {{ total }} 件商品
      </div>
    </div>

    <Comfirmbuy
      :product="comfirmproduct.product"
      :show-dialog="comfirmproduct.showComfirmBuyDialog"
      @close-dialog="dialogControler.closeDialog"
    />
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { runWithBackoff } from '~/composables/useBackoff'
import Comfirmbuy from '~/components/comfirmbuy.vue'
import ProductCard from '~/components/productCard.vue'
import { getProductList } from '~/composables/getProduct'
import { useComfirmBuyStore } from '~/utils/useProductStore'
import type { productResponse } from '~/types/product'
definePageMeta({ layout: 'home-page-layout' })
const showDialog = ref(false)
const comfirmproduct = useComfirmBuyStore()

const products = ref<productResponse[]>([])
const loading = ref(false)
const initialLoading = ref(true)
const finished = ref(false)
const total = ref(0)
const page = ref(1)
const pageSize = 9
// 重试计数：如果连续发生多次失败，停止重试以避免无限请求
const retryCount = ref(0)
const maxRetries = 3

class dialogControl {
  openDialog(product: productResponse) {
    comfirmproduct.setProduct(product)
    comfirmproduct.openComfirmBuyDialog()
    console.log('打开对话框，产品ID=', product.id)
  }
  closeDialog() {
    comfirmproduct.closeComfirmBuyDialog()
    comfirmproduct.resetProduct()
  }
}
const dialogControler = new dialogControl()

const loadMore = async () => {
  if (loading.value || finished.value) return
  loading.value = true
  try {
    // 使用通用的指数退避重试助手，runWithBackoff 会在内部处理重试与延迟
    const result = await runWithBackoff(() => getProductList({ page: page.value, pageSize }), {
      retries: 3,
      baseDelay: 600,
    })
    console.log('加载到的产品列表：', result)
    // result 肯定是有值或会被上层抛出
    products.value.push(...result.list)
    total.value = result.total
    finished.value = !result.hasMore
    if (result.hasMore) {
      page.value += 1
    }
  } catch (error) {
    // runWithBackoff 超出重试次数会抛出错误到这里
    console.error('加载产品列表最终失败：', error)
    finished.value = true // 停止自动再试，交由用户手动刷新或重试按钮
    ElMessage.error('加载产品列表失败，请检查网络或稍后刷新')
  } finally {
    loading.value = false
    initialLoading.value = false
  }
}

onMounted(async () => {
  await loadMore()
})
</script>

<style scoped>
.product-list-container {
  height: 890px;
  overflow-y: auto;
  padding: 16px;
  background-color: rgba(255, 255, 255, 0.244);
  border-radius: 12px;
  box-shadow: inset 0 0 8px rgba(0, 0, 0, 0.05); /* 内阴影，轻柔分层 */
}

/* 自定义滚动条（平面风） */
.product-list-container::-webkit-scrollbar {
  width: 8px;
}
.product-list-container::-webkit-scrollbar-track {
  background: transparent;
}
.product-list-container::-webkit-scrollbar-thumb {
  background-color: #cfd8dc;
  border-radius: 4px;
}
.product-list-container::-webkit-scrollbar-thumb:hover {
  background-color: #90a4ae;
}

.product-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
  list-style: none;
  margin: 0;
  padding: 0;
}

/* 每个商品卡片项 */
.product-item {
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}
.product-item:hover {
  transform: scale(1.01);
}
</style>
