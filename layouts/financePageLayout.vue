<!--
    financePageLayout.vue
    说明：应用财务页面的布局组件
    布局职责：
        - 渲染顶部导航栏（TopBar），并通过默认插槽在左侧渲染导航项
        - 提供右侧命名插槽 `end` 用于放置登录/用户菜单等操作项
        - 主内容通过默认 slot 插入（位于页面中间、带滚动）
-->
<template>
    <div class="flex-col h-screen ">
        <div class="h-16">
            <TopBar>
            <!-- 默认插槽内容（会出现在左侧） -->
                <template #default>
                  <Button icon="pi pi-home"  @click="navigateTo('/middlepage')" severity="secondary" raised variant="outlined"/>
                </template>
            <!-- 命名插槽 end：放置在右侧 -->
            <template #end>
                <div class="flex items-center gap-5">
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
            </template>
            </TopBar>
        </div>
    <div class="flex-col overflow-auto mt-1 ml-10 mr-10 items-center">
        <slot/>
    </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import TopBar from '~/components/topBar.vue'
import { menuData } from '~/config/topbarRoute'
import Menu from 'primevue/menu'
import Avatar from 'primevue/avatar'
const userStore = useUserStore()
const menu = ref()

// 在运行时把纯数据转换为 PrimeVue Menu 所需的 model（包含 command 函数）
const menuItems = computed(() =>
    menuData.map((m) => {
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

const toggleMenu = (event: Event) => {
  menu.value.toggle(event)
}
const activeIndex = ref(0)

</script>

<style scoped>
.nav-item {
    height: 100%;
    width: 6rem;
    display: flex;
    align-items: center;
    justify-items: center;
    cursor: pointer;
    user-select: none;
    transition: background 250ms ease, transform 150ms ease, box-shadow 250ms ease;
    background: transparent;
    color: inherit;
}

.nav-item:hover {
    background: linear-gradient( rgba(92, 92, 92, 0.591));
    transform: translateY(-1px);
}


</style>
