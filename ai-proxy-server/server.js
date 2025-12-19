/**
 * AI 代理服务器 (BFF - Backend For Frontend)
 * 功能：处理 Coze AI 聊天流式转发，解决跨域问题，保护 Token
 */

import express from 'express'
import cors from 'cors'
import dotenv from 'dotenv'

// 加载环境变量
dotenv.config()

const app = express()
const PORT = process.env.PORT || 15000

// 配置 CORS - 允许前端访问
const corsOptions = {
  origin: [
    'http://localhost:3000',      // Nuxt 开发服务器
    'http://localhost:5173',      // Vite 开发服务器
    'http://127.0.0.1:3000',
    'http://127.0.0.1:5173',
    // 生产环境可以添加你的前端域名
    // 'https://yourdomain.com'
  ],
  credentials: true,
  methods: ['GET', 'POST', 'OPTIONS'],
  allowedHeaders: ['Content-Type', 'Authorization']
}

app.use(cors(corsOptions))
app.use(express.json())

// 从环境变量获取配置（如果未设置，使用默认值）
const COZE_TOKEN = process.env.COZE_TOKEN || 'pat_npKKR94VWrHRCHbON5rHASqxVVSPk6Gq1qBu7M5Lp5JNONxB7jPc6WSGwcYdqbKu'
const COZE_BOT_ID = process.env.COZE_BOT_ID || '7585274475898486826'
const COZE_API_URL = 'https://api.coze.cn/v3/chat'

// 健康检查接口
app.get('/health', (req, res) => {
  res.json({ 
    status: 'ok', 
    message: 'AI Proxy Server is running',
    timestamp: new Date().toISOString()
  })
})

// 流式聊天接口
app.post('/api/stream-chat', async (req, res) => {
  try {
    const { message, user_id } = req.body

    // 验证请求参数
    if (!message) {
      return res.status(400).json({ 
        error: 'Message is required' 
      })
    }

    const userId = user_id || 'default_user'

    console.log('收到聊天请求:', {
      message: message.substring(0, 50) + '...',
      user_id: userId,
      timestamp: new Date().toISOString()
    })

    // 构建请求到 Coze API
    const requestBody = {
      bot_id: COZE_BOT_ID,
      user_id: userId,
      stream: true,
      additional_messages: [
        {
          role: 'user',
          content: message,
          content_type: 'text'
        }
      ]
    }

    console.log('转发请求到 Coze API:', {
      bot_id: COZE_BOT_ID,
      user_id: userId,
      stream: true
    })

    // 发起流式请求到 Coze API
    const response = await fetch(COZE_API_URL, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${COZE_TOKEN}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requestBody)
    })

    // 检查响应状态
    if (!response.ok) {
      const errorText = await response.text()
      console.error('Coze API 错误:', response.status, errorText)
      return res.status(response.status).json({ 
        error: errorText 
      })
    }

    // 检查响应体是否存在
    if (!response.body) {
      console.error('Coze API 响应体为空')
      return res.status(500).json({ 
        error: 'Response body is empty' 
      })
    }

    // 设置响应头，支持 SSE (Server-Sent Events)
    res.setHeader('Content-Type', 'text/event-stream')
    res.setHeader('Cache-Control', 'no-cache')
    res.setHeader('Connection', 'keep-alive')
    res.setHeader('X-Accel-Buffering', 'no') // 禁用 Nginx 缓冲

    console.log('开始流式传输...')
    console.log('响应头信息:', {
      'content-type': response.headers.get('content-type'),
      'transfer-encoding': response.headers.get('transfer-encoding')
    })

    // 获取流式数据的读取器
    const reader = response.body.getReader()
    const decoder = new TextDecoder('utf-8')

    let buffer = ''
    let chunkCount = 0
    let dataLineCount = 0

    console.log('开始读取流式数据...')

    try {
      // 循环读取数据流
      while (true) {
        const { done, value } = await reader.read()

        if (done) {
          console.log(`流式传输完成，共处理 ${chunkCount} 个数据块，${dataLineCount} 条数据行`)
          // 处理缓冲区中剩余的数据
          if (buffer.trim()) {
            const trimmedLine = buffer.trim()
            console.log('处理缓冲区剩余数据:', trimmedLine.substring(0, 100))
            if (trimmedLine.startsWith('data: ')) {
              res.write(trimmedLine + '\n\n')
            } else if (trimmedLine) {
              res.write(`data: ${trimmedLine}\n\n`)
            }
          }
          // 发送结束标志
          res.write('data: [DONE]\n\n')
          res.end()
          break
        }

        chunkCount++
        
        // 将二进制数据解码为文本
        const chunk = decoder.decode(value, { stream: true })
        buffer += chunk

        // 按行处理 SSE 格式数据
        const lines = buffer.split('\n')
        buffer = lines.pop() || '' // 保留最后一个不完整的行

        for (const line of lines) {
          const trimmedLine = line.trim()
          
          // 跳过空行
          if (!trimmedLine) continue

          // 处理 SSE 格式：data: {...}
          if (trimmedLine.startsWith('data: ')) {
            dataLineCount++
            // 直接转发数据行（保持 SSE 格式）
            res.write(trimmedLine + '\n\n')
            
            // 前 5 条数据输出日志用于调试
            if (dataLineCount <= 5) {
              console.log(`转发数据 #${dataLineCount}:`, trimmedLine.substring(0, 100))
            }
            
            // 刷新缓冲区，确保数据立即发送（如果支持）
            if (typeof res.flush === 'function') {
              res.flush()
            }
          } else {
            // 如果不是 data: 格式，也转发（可能是其他 SSE 事件）
            res.write(`data: ${trimmedLine}\n\n`)
            if (dataLineCount <= 5) {
              console.log(`转发非标准格式数据:`, trimmedLine.substring(0, 100))
            }
            if (typeof res.flush === 'function') {
              res.flush()
            }
          }
        }

        // 每处理 10 个数据块输出一次日志
        if (chunkCount % 10 === 0) {
          console.log(`已处理 ${chunkCount} 个数据块，缓冲区大小: ${buffer.length}`)
        }
      }
    } catch (streamError) {
      console.error('流式传输错误:', streamError)
      // 处理缓冲区中剩余的数据
      if (buffer.trim()) {
        res.write(`data: ${buffer.trim()}\n\n`)
      }
      res.write(`data: ${JSON.stringify({ error: 'Stream error', message: streamError.message })}\n\n`)
      res.end()
    } finally {
      // 确保释放读取器
      try {
        reader.releaseLock()
      } catch (e) {
        // 忽略释放错误
      }
    }

  } catch (error) {
    console.error('API 代理错误:', error)
    res.status(500).json({ 
      error: 'Internal Server Error',
      message: error.message 
    })
  }
})

// 启动服务器
app.listen(PORT, () => {
  console.log(`
╔══════════════════════════════════════════════════╗
║   AI Proxy Server (BFF) 已启动                  ║
╠══════════════════════════════════════════════════╣
║  端口: ${PORT.toString().padEnd(43)}║
║  健康检查: http://localhost:${PORT}/health${' '.repeat(20)}║
║  聊天接口: http://localhost:${PORT}/api/stream-chat${' '.repeat(7)}║
╚══════════════════════════════════════════════════╝
  `)
  console.log('等待前端请求...')
})

