/**
 * Coze AI 聊天 API 代理接口
 * 功能：将前端的请求转发到 Coze API，实现流式数据传输
 * 优势：解决 CORS 问题，保护 API Token 不被暴露在前端
 */

export default defineEventHandler(async (event) => {
  // 只允许 POST 请求
  if (event.node.req.method !== 'POST') {
    throw createError({
      statusCode: 405,
      statusMessage: 'Method Not Allowed'
    })
  }

  try {
    // 从 runtimeConfig 获取配置（Token 和 Bot ID 存储在服务器端）
    const config = useRuntimeConfig(event)
    const token = config.cozeToken || process.env.COZE_TOKEN || 'pat_npKKR94VWrHRCHbON5rHASqxVVSPk6Gq1qBu7M5Lp5JNONxB7jPc6WSGwcYdqbKu'
    const botId = config.cozeBotId || process.env.COZE_BOT_ID || '7585274475898486826'

    // 读取请求体
    const body = await readBody(event)
    const { message, user_id } = body

    if (!message) {
      throw createError({
        statusCode: 400,
        statusMessage: 'Message is required'
      })
    }

    // 构建请求到 Coze API
    const cozeUrl = 'https://api.coze.cn/v3/chat'
    
    console.log('转发请求到 Coze API:', {
      bot_id: botId,
      user_id: user_id || 'default_user',
      stream: true
    })

    // 使用 fetch 发起流式请求
    const response = await fetch(cozeUrl, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        bot_id: botId,
        user_id: user_id || 'default_user',
        stream: true,
        additional_messages: [
          {
            role: 'user',
            content: message,
            content_type: 'text'
          }
        ]
      })
    })

    // 检查响应状态
    if (!response.ok) {
      const errorText = await response.text()
      console.error('Coze API 错误:', response.status, errorText)
      throw createError({
        statusCode: response.status,
        statusMessage: errorText
      })
    }

    // 检查响应体是否存在
    if (!response.body) {
      throw createError({
        statusCode: 500,
        statusMessage: 'Response body is empty'
      })
    }

    // 设置响应头，支持 SSE (Server-Sent Events)
    setResponseHeader(event, 'Content-Type', 'text/event-stream')
    setResponseHeader(event, 'Cache-Control', 'no-cache')
    setResponseHeader(event, 'Connection', 'keep-alive')
    setResponseHeader(event, 'X-Accel-Buffering', 'no') // 禁用 Nginx 缓冲

    // 将 Coze 的流式响应直接管道传输到客户端
    // 使用 Node.js 的流式传输
    const reader = response.body.getReader()
    const decoder = new TextDecoder('utf-8')

    // 创建一个可读流用于传输数据
    const stream = new ReadableStream({
      async start(controller) {
        try {
          while (true) {
            const { done, value } = await reader.read()

            if (done) {
              console.log('流式传输完成')
              controller.close()
              break
            }

            // 将二进制数据解码并转发
            const chunk = decoder.decode(value, { stream: true })
            controller.enqueue(new TextEncoder().encode(chunk))
          }
        } catch (error: any) {
          console.error('流式传输错误:', error)
          controller.error(error)
        } finally {
          reader.releaseLock()
        }
      }
    })

    // 返回流式响应
    return stream

  } catch (error: any) {
    console.error('API 代理错误:', error)
    
    // 如果是 createError，直接抛出
    if (error.statusCode) {
      throw error
    }

    // 其他错误返回 500
    throw createError({
      statusCode: 500,
      statusMessage: error.message || 'Internal Server Error'
    })
  }
})

