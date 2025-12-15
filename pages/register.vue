<template>
    <section class="flex min-h-screen items-center justify-center bg-emerald-50/40 py-12 px-6">
        <Card class="w-full max-w-3xl rounded-3xl border border-white/50 bg-white/95 shadow-2xl backdrop-blur-md">
            <template #content>
                <div class="space-y-6">
                    <header class="space-y-2 text-center">
                        <h1 class="text-3xl font-semibold text-slate-900">创建您的账户</h1>
                        <p class="text-sm text-slate-500">注册后即可体验完整的智慧农业服务</p>
                    </header>
                    <RegisterForm class="mx-auto max-w-2xl" @login="handleLogin" @submit="submitRegister" />
                </div>
            </template>
        </Card>
    </section>
</template>

<script setup lang="ts">
import Card from 'primevue/card'
import { useToast } from 'primevue/usetoast'
import RegisterForm from '~/components/registerform.vue'
import { onRegister } from '~/composables/useRegister'
import type { RegisterRequest } from '~/types/register'

definePageMeta({ layout: 'home-page-layout' })

const toast = useToast()

const handleLogin = () => {
    navigateTo('/login')
}

async function submitRegister(payload: RegisterRequest) {
    try {
        const res = await onRegister(payload)
        console.log('后端返回', res)
        toast.add({ severity: 'success', summary: '注册成功', detail: '请登录以继续', life: 2500 })
        navigateTo('/login')
    } catch (err) {
        toast.add({ severity: 'error', summary: '注册失败', detail: '请稍后重试', life: 2500 })
        console.error('注册失败', err)
    }
}
</script>

<style scoped>
</style>
