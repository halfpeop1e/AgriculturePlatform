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
    <ul v-infinite-scroll="load"  style="overflow: auto;height: 890px;">
    <el-card v-for="i in count" :key="i" class="h-40">{{ i }}</el-card>
  </ul>
</template>

<script setup lang="ts">
definePageMeta({ layout: 'home-page-layout' })
const loading=false
const count = ref(10)
const load = () => {
      count.value += 2
}
</script>

<style scoped>

</style>