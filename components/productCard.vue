<!--
  productCard.vue
  说明：产品展示卡片，包含图片、简介、价格、库存与购买入口
  Props:
    - product: 产品对象（id, name, image, description, price, stock, saler）
  Emits:
    - openDialog: 点击购买时触发，父组件可打开购买弹窗并传入 product
  注意点：图片使用 object-cover 保持封面显示，购买按钮触发 openDialog 事件。
-->
<template>
  <el-card class="w-full hover:shadow-lg transition-shadow duration-300 ease-in-out">
    <div class="flex gap-6">
      <!-- 左侧产品图片 -->
      <div class="w-48 aspect-[4/3] overflow-hidden rounded-lg">
        <imageShow :images="product.image"/>
      </div>

      <!-- 右侧信息 -->
      <div class="flex flex-col justify-between flex-1">
        <!-- 产品名称 -->
        <div class="text-xl font-semibold text-gray-800 mb-1">
          {{ product.name }}
        </div>

        <!-- 产品描述 -->
        <div class="text-gray-500 text-sm flex-1 leading-relaxed">
          {{ product.description }}
        </div>

        <!-- 底部价格与库存 -->
        <div class="flex justify-between items-center mt-3">
            <div class="flex items-baseline gap-1">
            <div class="text-lg font-bold text-green-600">
            ￥{{ product.price }}
          </div>
                <div class="text-lg">/件</div>
            </div>
          
          <div class="text-sm text-gray-400">
            剩余：{{ product.stock }} 件
          </div>
        </div>
         
        <div class="flex justify-end mt-2">
          <div class=" text-xl flex-1 leading-relaxed">
         商户：{{ product.saler }}
          </div>
          <div v-if="!routedivided(route.path)">
            <el-button type="success" size="default" @click="openBuyDialog">购买</el-button>
          </div>
          <div class="flex" v-if="routedivided(route.path)">
          <el-button type="primary" size="default" @click="openBuyDialog">编辑</el-button>
          <el-button type="danger" size="default" @click="openBuyDialog">删除</el-button>
          </div>
        </div>
      </div>
    </div>
  </el-card>

  

</template>

<script setup lang="ts">
const route = useRoute()
interface Product {
  id: string
  name: string
  image:{ url: string}[]
  description: string
  price: number
  stock: number
  saler: string
}

const props = defineProps<{
  product: Product
}>()


function openBuyDialog() {
  emit('openDialog')
}
const emit = defineEmits<{
    (e: 'openDialog'): void
  }>()
function routedivided(route:string){
    const r=(route || '').toLowerCase()
    if(r.includes('myrelease')){
      return true
    }
    return false
}
</script>
<style scoped>
.aspect-\[4\/3\] {
  aspect-ratio: 4 / 3;
}
</style>
