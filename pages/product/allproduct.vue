<template>
    <div v-if="loading">
    <div class="m-10"></div>
    <el-skeleton :rows="4" animated />
    <div class="m-10"></div>
    <el-skeleton :rows="4" animated />
    <div class="m-10"></div>
    <el-skeleton :rows="4" animated />
    <div class="m-10"></div>
    <el-skeleton :rows="4" animated />      
    </div>
    <div class="product-list-container">
    <ul v-infinite-scroll="load" class="product-list">
      <li
        v-for="(item, index) in products"
        :key="index"
        class="product-item"
      >
        <ProductCard :product="item" />
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import ProductCard from '~/components/productCard.vue'
import { getProductList } from '~/composables/getProduct'
definePageMeta({ layout: 'home-page-layout' })
const loading=false
const productslist=await getProductList()
const products = ref([
  {
    name: '绿色有机苹果',
    image: '/images/apple.jpg',
    description: '源自无污染果园，新鲜香甜可口，富含维生素C。',
    price: 12.8,
    stock: 56,
    saler: '果园直供'
  },
  {
    name: '土鸡蛋',
    image: '/images/egg.jpg',
    description: '农家散养土鸡，新鲜营养不打药。',
    price: 8.5,
    stock: 120,
    saler: '农家直供'
  },
  {
    name: '香甜玉米',
    image: '/images/corn.jpg',
    description: '现摘玉米粒饱满，香糯可口。',
    price: 3.5,
    stock: 300,
    saler: '田园直供'
  }
])

// 模拟懒加载
const load = () => {
  setTimeout(() => {
    products.value.push(
      ...Array(3)
        .fill(0)
        .map((_, i) => ({
          name: `助农产品 ${products.value.length + i + 1}`,
          image: '/testproductimg.jpg',
          description: '助农优质产品，品质保障。',
          price: 10 + Math.random() * 5,
          stock: 50 + Math.floor(Math.random() * 100),
          saler: '助农直供'
        }))
    )
  }, 800)
}
</script>

<style scoped>
.product-list-container {
  height: 890px;
  overflow-y: auto;
  padding: 16px;
  background-color: #f9fafb; /* 淡灰底，显平面感 */
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