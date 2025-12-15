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
          <ProductCard :product="item" @open-buy-dialog="dialogControler.openDialog(item)" @delete-product="handleDelete(item.id)" @open-edit-dialog="dialogControler.openEditDialog(item)" />
        </li>
      </ul>

      <el-empty
        v-if="!products.length && !loading"
        description="暂时没有您发布的助农产品"
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
      <ProductEditDialog 
      :model-value="editproduct.showEditProductDialog" 
      :product="editproduct.product" 
      @save="handleSaveEdit"
      @update:modelValue="val => editproduct.showEditProductDialog = val" @cancel="dialogControler.closeEditDialog" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { runWithBackoff } from '~/composables/useBackoff'
import ProductCard from '~/components/productCard.vue'
import { getProductList } from '~/composables/getProduct'
import { EditMyProduct } from '~/composables/useProduct'
import { useComfirmBuyStore , useEditProductStore} from '~/utils/useProductStore'
import type { productResponse } from '~/types/product'
definePageMeta({ layout: 'home-page-layout' })
const showDialog = ref(false)
const comfirmproduct = useComfirmBuyStore()
const editproduct = useEditProductStore()
const products = ref<productResponse[]>([])
const loading = ref(false)
const initialLoading = ref(true)
const finished = ref(false)
const total = ref(0)
const page = ref(1)
const pageSize = 9

class dialogControl {
  openDialog(product: productResponse) {
    comfirmproduct.setProduct(product)
    showDialog.value = true
  }
  closeDialog() {
    showDialog.value = false
    comfirmproduct.resetProduct()
  }
  openEditDialog(product: productResponse) {
    console.log('Opening edit dialog for product:', product)
    editproduct.setProduct(product)
    editproduct.openEditProductDialog()
  }
  closeEditDialog() {
    editproduct.resetProduct()
    editproduct.closeEditProductDialog()
  }
}
const dialogControler = new dialogControl()

const loadMore = async () => {
  if (loading.value || finished.value) return
  loading.value = true
  const userStore = useUserStore()
  try {
    // 使用通用的指数退避重试助手，runWithBackoff 会在内部处理重试与延迟
    const result = await runWithBackoff(() => getProductList({ page: page.value, pageSize,salerId:userStore.userId }), {
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
async function handleSaveEdit(editedData: any){
  try{
    await EditMyProduct(editedData)
    ElMessage.success('商品信息已更新')
    // 刷新商品列表或更新本地状态
    products.value = []
    page.value = 1
    finished.value = false
    await loadMore()
  }catch(err){
    ElMessage.error('更新商品信息失败')
  }
}
async function handleDelete(id:string){
  try{
      const response= await DeleteMyProduct(id)
      if(response?.status==200){
        loadMore()
        return
      }
      else{
        throw new Error('删除商品失败')
      }
  }
  catch(err){
      ElMessage.error('删除商品失败，请稍后重试')
  }
      loadMore()
}
onMounted(async () => {
  await loadMore()
})
</script>

<style scoped>

</style>
