export default defineNuxtRouteMiddleware((to) => {
    // 允许访问的公开路径
    const publicPaths = ['/','/homepage', '/login', '/register']

    // 如果是公开路径，直接放行
    if (publicPaths.includes(to.path)) return

    // 检查是否已登录（通过 cookie 中的 AuthToken）
    const authToken = useCookie('AuthToken')?.value

    // 如果没有登录凭证，重定向到登录页
    if (!authToken) return navigateTo('/login')
})
