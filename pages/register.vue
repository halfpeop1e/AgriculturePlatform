<template>
    <div class="flex items-center justify-center h-screen">
        <RegisterForm class="mb-20" @login="handleLogin"/>
    </div>
</template>

<script setup lang="ts">

import RegisterForm from '~/components/registerform.vue';

definePageMeta({ layout: 'home-page-layout' })

const handleLogin=()=>{
    navigateTo('/login')
}
</script>

<style scoped>
</style>