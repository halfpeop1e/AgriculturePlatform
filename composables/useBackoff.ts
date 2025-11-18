/*
  useBackoff.ts
  说明：提供一个通用的指数退避/重试助手 runWithBackoff
  用法：
    import { runWithBackoff } from '~/composables/useBackoff'
    const result = await runWithBackoff(() => apiCall(), { retries: 3, baseDelay: 500 })
  返回：如果在重试次数内成功返回值，失败则抛出最后一个错误
*/

export interface BackoffOptions {
  retries?: number // 最大重试次数（不含初始尝试），默认为 3
  baseDelay?: number // 基础延迟毫秒，指数增长，默认为 500
}

export function sleep(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

export async function runWithBackoff<T>(fn: () => Promise<T>, options: BackoffOptions = {}): Promise<T> {
  const { retries = 3, baseDelay = 500 } = options
  let attempt = 0
  while (true) {
    try {
      return await fn()
    } catch (err) {
      attempt += 1
      if (attempt > retries) {
        // 所有尝试均失败，抛出最后的错误
        throw err
      }
      // 指数退避：baseDelay * 2^(attempt-1)
      const delay = Math.round(baseDelay * Math.pow(2, attempt - 1))
      // 轻度抖动，避免群体重试冲突
      const jitter = Math.round(Math.random() * Math.min(300, delay))
      await sleep(delay + jitter)
    }
  }
}
