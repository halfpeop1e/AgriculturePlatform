<template>
  <el-card class="w-full hover:shadow-lg transition-shadow duration-300 ease-in-out">
    <div class="flex gap-6">
      <!-- 左侧产品图片 -->
      <div class="w-48 aspect-[4/3] overflow-hidden rounded-lg">
        <img
          :src="product.image"
          alt="product"
          class="w-full h-full object-cover hover:scale-105 transition-transform duration-300 ease-in-out cursor-pointer"
        />
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
         <div class=" text-xl flex-1 leading-relaxed">
         商户：{{ product.saler }}
        </div>
      </div>
    </div>
  </el-card>
</template>

<script setup lang="ts">
interface Product {
  name: string
  image: string
  description: string
  price: number
  stock: number
  saler: string
}

const props = defineProps<{
  product: Product
}>()
</script>

<style scoped>
.aspect-\[4\/3\] {
  aspect-ratio: 4 / 3;
}
</style>
