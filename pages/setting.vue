<template>
  <div class="p-6">
    <div class="relative mx-auto flex max-w-6xl gap-6">
      <div class=" left-6 top-20 z-20">
        <Button icon="pi pi-bars" outlined rounded @click="toggleSidebar" />
      </div>

      <SizeBar :nav="settingNav" />

      <Card class="flex-1 rounded-2xl border border-slate-200 bg-white shadow-lg">
        <template #content>
          <div class="min-h-[540px] rounded-2xl bg-slate-50/60 p-6">
            <router-view />
          </div>
        </template>
      </Card>
    </div>
  </div>
</template>

<script setup lang="ts">
import Button from 'primevue/button'
import Card from 'primevue/card'
import SizeBar from '~/components/sizeBar.vue'
import { settingNav } from '~/config/settingRoute'


const buttonstore = buttonStore()
const router = useRouter()

function toggleSidebar() {
  buttonstore.isvisible = !buttonstore.isvisible
}

onMounted(() => {
  if (router.currentRoute.value.path === '/setting') {
    router.replace('/setting/profile')
  }
})
</script>

<style scoped>
</style>
