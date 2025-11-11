<template>
    <div class="flex items-center justify-center h-screen">
        <LoginForm class="mb-20" @register="handleRegister" @setactivetime="SetActiveTime"/>
    </div>
</template>

<script setup lang="ts">
import LoginForm from '~/components/loginForm.vue';
import moment from 'moment';
definePageMeta({ layout: 'home-page-layout' })

const handleRegister=()=>{
    navigateTo('/register')
}
function SetActiveTime(){
    const UserStore=useUserStore()
    UserStore.lastActive=moment().format('YYYY-MM-DD HH:mm:ss')
}

</script>

<style scoped>
</style>
