import type { LayoutKey } from "~/.nuxt/types/layouts"

const goBack = () => {
  window.history.back()
}
export default defineNuxtRouteMiddleware((to) => {
  if(to.path=='/middlepage'){
    goBack()
  }
  const ispreProfile=to.path.startsWith('/profile')
  if(ispreProfile&&import.meta.client){
    const userStore=useUserStore()
    const userRole = userStore.role
    const layoutMap=ref<LayoutKey>('profile-page-layout')
    if(userRole==='expert'||userRole==='finance'){
      layoutMap.value='finance-page-layout' as LayoutKey
    }
    else{
      layoutMap.value='profile-page-layout' as LayoutKey
    }
     to.meta.layout=layoutMap.value
  }
})
