<template>
    <section class="flex min-h-screen items-center justify-center bg-slate-900/5 py-12 px-6">
        <Card class="w-full max-w-2xl rounded-3xl border border-white/40 bg-white/90 shadow-2xl backdrop-blur-md">
            <template #content>
                <div class="space-y-6">
                    <header class="space-y-2 text-center">
                        <h1 class="text-3xl font-semibold text-slate-900">欢迎回来</h1>
                        <p class="text-sm text-slate-500">登录账号以管理您的农业数字化业务</p>
                    </header>
                    <LoginForm class="mx-auto max-w-lg" @register="handleRegister" @setactivetime="SetActiveTime" />
                </div>
            </template>
        </Card>
    </section>
</template>

<script setup lang="ts">
import Card from 'primevue/card'
import LoginForm from '~/components/loginForm.vue'
import moment from 'moment'

definePageMeta({ layout: 'home-page-layout' })

const handleRegister = () => {
    navigateTo('/register')
}

function SetActiveTime() {
    const UserStore = useUserStore()
    UserStore.lastActive = moment().format('YYYY-MM-DD HH:mm:ss')
}
</script>

<style scoped>
</style>
