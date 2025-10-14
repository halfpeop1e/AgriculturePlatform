<template>
    <div class="flex items-center justify-center h-screen">
        <RegisterForm class="mb-20" @login="handleLogin" @submit="submitRegister"/>
    </div>
</template>

<script setup lang="ts">

import RegisterForm from '~/components/registerform.vue';
import { onRegister } from '~/composables/useRegister';
import type { RegisterRequest } from '~/types/register';
definePageMeta({ layout: 'home-page-layout' })

const handleLogin=()=>{
    navigateTo('/login')
}
async function submitRegister(payload:RegisterRequest){
    try{
       const res= await onRegister(payload)
       console.log('后端返回',res)
       navigateTo('/login')
    }catch(err){
        ElMessage.error('注册失败，请重试')
        console.error('注册失败',err)
    }
}
</script>

<style scoped>
</style>