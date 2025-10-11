<template>
    <div class="flex items-center justify-center h-screen">
        <LoginForm class="mb-20"/>
    </div>
</template>

<script setup lang="ts">
import LoginForm from '~/components/loginForm.vue';

definePageMeta({ layout: 'home-page-layout' })

</script>

<style scoped>
</style>