<!--
    homePageLayout.vue
    说明：应用首页的布局组件
    布局职责：
        - 渲染顶部导航栏（TopBar），并通过默认插槽在左侧渲染导航项
        - 提供右侧命名插槽 `end` 用于放置登录/用户菜单等操作项
        - 主内容通过默认 slot 插入（位于页面中间、带滚动）
    设计要点：TopBar 与 PrimeVue/Element 组件混用；导航数据来自 `~/config/topbarRoute`。
    注意：组件内部会把纯数据 menuData 映射成 PrimeVue Menu 所需的 model，command 使用 navigateTo 进行路由跳转。
-->
<template>
    <div class="flex-col h-screen ">
        <div class="h-16">
            <TopBar>
            <!-- 默认插槽内容（会出现在左侧） -->
                <template #default>
                    <div class="text-white text-xl flex h-full items-center">
                                <el-dropdown class="flex items-center justify-center"
                                    v-for="(item, i) in nav"
                                    :key="i"
                                    @click="setActive(i)"
                                    @keydown.enter="setActive(i)"
                                    tabindex="0"
                                    size="large"
                                    :class="[
                                        'nav-item',
                                        activeIndex === i ? 'active' : ''
                                    ]"
                                >
                            <span class="h-full w-full flex items-center justify-center text-lg">{{ item.label }}</span>
                            <template #dropdown>
                                <el-dropdown-menu>
                                    <el-dropdown-item
                                        v-for="(subitem, i) in item.items"
                                        :key="i"
                                        @click="handleClick(subitem.routername)"
                                        >
                                        {{ subitem.label }}
                                    </el-dropdown-item>
                                </el-dropdown-menu>
                            </template>
                        </el-dropdown>
                    </div>
                </template>
            <!-- 命名插槽 end：放置在右侧 -->
            <template #end>
                <div class="flex items-center gap-5">
                   <el-button v-if="!userStore.islogin" type="success" @click="navigateTo('/login')">登录 ｜ 注册</el-button>
                  <Avatar
      image="https://avatars.githubusercontent.com/u/9919?s=280&v=4"
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
    <div class="flex-col overflow-auto mt-1 ml-20 mr-20 items-center">
        <slot/>
    </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import TopBar from '~/components/topBar.vue'
import { nav, menuData } from '~/config/topbarRoute'
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
                command: () => console.log('退出登录')
            }
        }
        return {
            label: m.label,
            icon: m.icon,
            command: () => navigateTo((m as any).to)
        }
    })
)
const handleClick = (routerlink:string) => {
  navigateTo(routerlink)
}
const toggleMenu = (event: Event) => {
  menu.value.toggle(event)
}
const activeIndex = ref(0)

function setActive(i: number) {
    activeIndex.value = i
}
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
