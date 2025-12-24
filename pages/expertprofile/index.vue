<template>
  <div>

  </div>
</template>

<script setup lang="ts">


onMounted(()=>{ 
  const userStore = useUserStore()
  navigateTo(`/expertprofile/${userStore.userId}`)
})
</script>

<style scoped>

</style>
