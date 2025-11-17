/*
  userLogin.global.ts
  说明：全局路由中间件，用于保护需要登录的页面
  行为：
    - 允许访问的公开路径：'/', '/homePage', '/login', '/register'
    - 若访问非公开路径，检查 cookie 中的 `AuthToken`（优先 useCookie）
    - 若未检测到 token，则显示提示并跳转到 `/login`
  注意：当前实现只做“存在性检查”，未校验 token 有效性或过期；可以扩展为调用后端验证或解析 JWT 检查 exp
*/
export default defineNuxtRouteMiddleware((to) => {
    // 允许访问的公开路径
    const publicPaths = ['/','/homePage', '/login', '/register','/css','/setting/profile','/product/sellproduct']
    const userStore = useUserStore()
    // 如果是公开路径，直接放行
    if (publicPaths.includes(to.path)) {
      return
    }
    // 检查是否已登录（通过 cookie 中的 AuthToken）
    const authToken = useCookie('AuthToken')?.value

    // 如果没有登录凭证，重定向到登录页
    if (!authToken&&!userStore.tokens) {
      return navigateTo('/login')
    }
})
