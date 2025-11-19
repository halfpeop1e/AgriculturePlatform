// @ts-ignore - Nuxt runtime types may not be available in editor linting environment
import { defineNuxtPlugin } from '#app'
import { useUserStore } from '~/utils/userStore'
import { getUserProfile } from '~/composables/getProfile'

// 客户端启动时尝试从 cookie 中恢复用户会话：
// - 读取 useCookie('AuthToken')，若存在则解析 token payload 提取 userId
// - 将 token 与 userId 写入 userStore 并调用 getUserProfile() 恢复用户资料
// - 若解析或请求失败，清理 cookie 与本地 token

export default defineNuxtPlugin((nuxtApp: any) => {
  // 仅在浏览器端执行
  if (process.client) {
    const tryRestore = async () => {
      try {
  // useCookie 在 Nuxt 客户端可用
  // 将 cookie 视为 any 以避免类型噪音（编辑器中未必能解析 Nuxt 全局类型）
  // @ts-ignore
  const cookie: any = useCookie('AuthToken')
  const token = cookie?.value as string | undefined
        if (!token) return

        const userStore = useUserStore()

        // 若 store 已经标记为登录则不重复恢复
        if (userStore.islogin) return

        // 简单解析 JWT payload（base64url）并尝试从常见字段读取 userId
        const parts = token.split('.')
        if (parts.length < 2) {
          throw new Error('非法的 token 格式')
        }

  const payloadB64 = parts[1] as string
        // base64url -> base64
        const base64 = payloadB64.replace(/-/g, '+').replace(/_/g, '/')
        // atob -> 解码，再 decodeURIComponent to preserve utf8
        const jsonStr = decodeURIComponent(
          Array.prototype.map
            .call(window.atob(base64), (c: string) => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2))
            .join('')
        )
        const payload = JSON.parse(jsonStr || '{}')

        // 支持多种可能的字段名
        const userId =
          payload.userId || payload.user_id || payload.uid || payload.id || payload.sub || (payload.data && (payload.data.userId || payload.data.id))

        if (userId) {
          userStore.userId = String(userId)
        }

        // 将 token 写入 store（后续请求的拦截器会从 cookie/localStorage/read 处读取；这里保持同步）
        userStore.tokens = token
        try {
          // 将 token 也写一份到 localStorage 以防拦截器不同位置读取
          localStorage.setItem('AuthToken', token)
        } catch (e) {
          // ignore
        }

        // 请求用户资料（受保护接口），若成功则 setUserProfile
        await getUserProfile()
        userStore.LoginSet()
        console.log('已从 cookie 恢复用户会话',userStore)
      } catch (err) {
        console.error('从 cookie 恢复用户会话失败：', err)
        try {
          // 清理可能的残留 token
          // @ts-ignore
          const cookie = useCookie('AuthToken')
          if (cookie) cookie.value = null
        } catch (e) {}
        try { localStorage.removeItem('AuthToken') } catch (e) {}
      }
    }

    // 在 micro task 中异步执行恢复（确保 Pinia 已初始化）
    setTimeout(() => {
      void tryRestore()
    }, 0)
  }
})
