<template>
  <div class="min-h-screen bg-slate-100">
    <header class="border-b border-slate-200 bg-white shadow-sm">
      <div class="mx-auto flex w-full max-w-6xl items-center justify-between px-6 py-4">
        <div class="flex items-center gap-4">
          <div class="flex h-12 w-12 items-center justify-center rounded-full bg-emerald-500 text-lg font-semibold text-white">
            QA
          </div>
          <div>
            <h1 class="text-xl font-semibold text-slate-900">专家后台</h1>
            <p class="text-sm text-slate-500">管理农业问答与专家解答流程</p>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <Avatar
            :image="userStore.avatar"
            shape="circle"
            size="large"
            class="cursor-pointer border border-gray-300 hover:ring-2 hover:ring-gray-300 transition-all"
          @click="toggleMenu"
          />
          <Menu
      ref="menu"
      :model="menuItems"
      :popup="true"
      style="margin-top: 10px;margin-left: 10px;"
    />
        </div>
      </div>
    </header>

    <main class="mx-auto w-full max-w-6xl px-6 py-10">
      <slot />
    </main>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import Avatar from 'primevue/avatar'
import { useUserStore } from '~/utils/userStore'
import {  menuExpertData } from '~/config/topbarRoute'

const userStore = useUserStore()
const menu = ref()
const defaultAvatar = 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png'

const toggleMenu = (event: Event) => {
  menu.value.toggle(event)
}
const menuItems = computed(() =>
    menuExpertData.map((m) => {
        if ((m as any).separator) return { separator: true }
        if ((m as any).action === 'logout') {
            return {
                label: m.label,
                icon: m.icon,
                command: () => {
                  logoutUser()
                  navigateTo('/homePage')
                }
            }
        }
        return {
            label: m.label,
            icon: m.icon,
            command: () => navigateTo((m as any).to)
        }
    })
)
</script>
