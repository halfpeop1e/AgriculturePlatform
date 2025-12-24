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
    const publicPaths = ['/','/homePage', '/login', '/register']
    const userStore = useUserStore()
    const isPublicPath = publicPaths.includes(to.path)

    // 检查是否已登录（通过 cookie 中的 AuthToken）
    const authToken = useCookie('AuthToken')?.value
    const isAuthenticated = Boolean(authToken || userStore.tokens)
    const isSettingRoute = to.path.startsWith('/setting')
    if (!isAuthenticated) {
      if (isPublicPath) {
        return
      }
      return navigateTo('/login')
    }
    if(isAuthenticated&&userStore.islogin){
      if(to.path=='/login'||to.path=='/register'){
        ElMessage.success("您已经登录")
        return navigateTo('/homePage')
       
      }
    }
    if (userStore.role === 'expert') {
      const completionPath = '/infomationcomplete'

      // 未完善资料：强制跳转到完善页，仅允许访问完善页本身
      if (!userStore.expertProfile?.name) {
        console.log(userStore.expertProfile?.name)
        if (to.path !== completionPath){
          console.log('跳转到专家资料完善页')
          return navigateTo(completionPath)
        } 
        return
      }
      // 允许的专家路由：以 /specialBoard 开头 或以 /expertprofile 开头
      const expertDashboardPath = '/specialBoard'
      const profilePathPrefix = '/expertprofile'
      const isExpertDashboard = to.path === expertDashboardPath || to.path.startsWith(expertDashboardPath)
      const isProfileRoute = to.path === profilePathPrefix || to.path.startsWith(profilePathPrefix)

      // 非以上允许路由时重定向到专家后台主页
      if (!isExpertDashboard && !isProfileRoute) {
        console.log('跳转到专家界面,topath', to.path)
        return navigateTo(expertDashboardPath)
      }
    
      return
    }

    if (userStore.role === 'finance') {
      const financeDashboardPath = '/finaceJustice'
      const profilePathPrefix = '/profile'
      const isProfileRoute = to.path.startsWith(profilePathPrefix)
      const isFinanceDashboard = to.path === financeDashboardPath

      if (!isFinanceDashboard && !isProfileRoute && !isSettingRoute) {
        console.log('跳转到财务界面')
        return navigateTo(financeDashboardPath)
      }
      return
    }

    if (isPublicPath) {
      console.log('访问公开路径，无需跳转')
      return
    }
})
