<template>
  <!-- 只在登录后显示 -->
  <div v-if="isLoggedIn" class="assistant-container">
    <!-- 悬浮球 -->
    <div 
      class="floating-ball"
      @click="toggleWindow"
      :class="{ 'ball-active': isWindowOpen }"
    >
      <span class="ball-text">智能助手</span>
    </div>

    <!-- 悬浮窗 -->
    <Transition name="window-fade">
      <div v-if="isWindowOpen" class="floating-window">
        <!-- 悬浮窗头部 -->
        <div class="window-header">
          <h3 class="window-title">智能助手</h3>
          <button class="close-btn" @click="closeWindow" aria-label="关闭">
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M15 5L5 15M5 5L15 15" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
            </svg>
          </button>
        </div>
        <!-- AI 对话内容区域 -->
        <div class="window-content">
          <!-- 消息列表区域 -->
          <div ref="messagesContainer" class="messages-container">
            <div v-if="messages.length === 0" class="welcome-message">
              <div class="welcome-icon">🤖</div>
              <p class="welcome-text">您好！我是智能助手，有什么可以帮助您的吗？</p>
            </div>
            <div
              v-for="(message, index) in messages"
              :key="index"
              :class="['message-item', message.role === 'user' ? 'user-message' : 'ai-message']"
            >
              <div class="message-avatar">
                <span v-if="message.role === 'user'">👤</span>
                <span v-else>🤖</span>
              </div>
              <div class="message-content">
                <div class="message-bubble">
                  <p class="message-text">{{ message.content }}</p>
                  <span class="message-time">{{ formatTime(message.timestamp) }}</span>
                </div>
              </div>
            </div>
            <!-- 加载指示器 -->
            <div v-if="isLoading" class="message-item ai-message">
              <div class="message-avatar">🤖</div>
              <div class="message-content">
                <div class="message-bubble">
                  <div class="typing-indicator">
                    <span></span>
                    <span></span>
                    <span></span>
                  </div>
                </div>
              </div>
            </div>
            <!-- 调试信息（开发时可见） -->
            <div v-if="debugInfo" class="debug-info">
              <details>
                <summary>调试信息（点击查看）</summary>
                <pre>{{ debugInfo }}</pre>
              </details>
            </div>
          </div>
          
          <!-- 输入区域 -->
          <div class="input-container">
            <div class="input-wrapper">
              <textarea
                v-model="inputText"
                @keydown.enter.exact.prevent="sendMessage"
                @keydown.shift.enter.exact="inputText += '\n'"
                placeholder="输入您的问题..."
                class="message-input"
                rows="1"
                ref="inputRef"
              ></textarea>
              <button
                @click="sendMessage"
                :disabled="!inputText.trim() || isLoading"
                class="send-button"
                :class="{ 'disabled': !inputText.trim() || isLoading }"
              >
                <svg v-if="!isLoading" width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M18 2L9 11M18 2L12 18L9 11M18 2L2 8L9 11" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
                <div v-else class="send-spinner"></div>
              </button>
            </div>
          </div>
        </div>
      </div>
    </Transition>

    <!-- 点击外部区域关闭悬浮窗 -->
    <div 
      v-if="isWindowOpen" 
      class="window-overlay"
      @click="closeWindow"
    ></div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch, nextTick } from 'vue'
import { storeToRefs } from 'pinia'

const userStore = useUserStore()
const { islogin } = storeToRefs(userStore)
const isLoggedIn = computed(() => islogin.value)

const isWindowOpen = ref(false)
const inputText = ref('')
const messages = ref<Array<{ role: 'user' | 'ai', content: string, timestamp: Date }>>([])
const isLoading = ref(false)
const messagesContainer = ref<HTMLElement | null>(null)
const inputRef = ref<HTMLTextAreaElement | null>(null)
const chatClient = ref<any>(null)
const chatId = ref<string>('') // 用于存储会话 ID
const debugInfo = ref<string>('') // 调试信息

// 格式化时间
const formatTime = (date: Date): string => {
  const hours = date.getHours().toString().padStart(2, '0')
  const minutes = date.getMinutes().toString().padStart(2, '0')
  return `${hours}:${minutes}`
}

// 滚动到底部
const scrollToBottom = () => {
  nextTick(() => {
    if (messagesContainer.value) {
      messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
    }
  })
}

// 创建或获取会话 ID
const getOrCreateChatId = async (): Promise<string> => {
  if (chatId.value) {
    return chatId.value
  }
  
  // 如果没有 chat_id，创建一个新的会话
  // 这里可以调用 Coze API 创建会话，或者使用一个简单的生成方式
  const newChatId = `chat_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`
  chatId.value = newChatId
  return newChatId
}

// 调用 API 代理（流式模式）
// 现在 Token 和 Bot ID 由服务器端处理，前端不再暴露敏感信息
const callCozeAPIStream = async (
  message: string,
  onChunk: (chunk: string) => void,
  onComplete: () => void,
  onError: (error: string) => void
): Promise<void> => {
  const userId = userStore.userId || userStore.userinfo.nickname || `user_${Date.now()}`
  
  // 获取 API 基础 URL（支持远程后端服务器或独立 BFF）
  const config = useRuntimeConfig()
  // 强制转换为字符串，避免 config.public.apiBaseUrl 被推断为 `{}` 导致调用 `replace` 报错
  const apiBaseUrl = String(config.public?.apiBaseUrl ?? '')
  
  // 构建 API 端点
  // 如果配置了远程服务器（如独立 BFF），使用完整 URL；否则使用相对路径（Nuxt Server API）
  const endpoint = apiBaseUrl 
    ? `${apiBaseUrl.replace(/\/$/, '')}/api/stream-chat`  // 独立 BFF 使用 /api/stream-chat
    : '/api/chat'  // Nuxt Server API 使用 /api/chat
  
  try {
    console.log('调用本地 API 代理:', endpoint)
    console.log('请求参数:', {
      message,
      user_id: userId
    })
    
    const response = await fetch(endpoint, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        message,
        user_id: userId
      })
    })

      console.log('API 响应状态:', response.status, response.statusText)

      if (!response.ok) {
        const errorText = await response.text()
        console.error('API 错误响应:', errorText)
        
        try {
          const errorData = JSON.parse(errorText)
          console.error('API 错误详情:', errorData)
          onError(`服务器错误: ${errorData.message || errorData.statusMessage || errorText}`)
        } catch (e) {
          onError(`请求失败: ${response.status} - ${errorText}`)
        }
        return
      }

      // 检查响应体是否存在
      if (!response.body) {
        onError('响应体为空，无法读取流数据')
        return
      }

      // 开始读取流式数据
      const reader = response.body.getReader()
      const decoder = new TextDecoder('utf-8')
      let buffer = ''
      
      // ✅ 核心修复：唯一的、不被清空的全量变量
      let fullResponseText = ''
      // ✅ 核心修复：记录当前事件类型，用于区分 delta 和 completed
      let currentEvent = ''

      console.log('开始读取流式数据...')

      try {
        while (true) {
          const { done, value } = await reader.read()

          if (done) {
            console.log('流式数据读取完成')
            break
          }

          // 将二进制数据解码为文本
          const chunk = decoder.decode(value, { stream: true })
          buffer += chunk

          // 按行分割处理 SSE 格式数据
          const lines = buffer.split('\n')
          buffer = lines.pop() || '' // 保留最后一个不完整的行（处理 JSON 截断）

          for (const line of lines) {
            const trimmedLine = line.trim()
            
            // 跳过空行
            if (!trimmedLine) continue
            
            // ✅ 核心修复：1. 记录当前行对应的事件类型
            if (trimmedLine.startsWith('event:')) {
              currentEvent = trimmedLine.replace('event:', '').trim()
              continue
            }
            
            // 2. 处理数据行
            if (!trimmedLine.startsWith('data:')) continue

            // 提取 data: 后面的 JSON 字符串
            const jsonStr = trimmedLine.substring(5).trim()
            
            // 检查结束标志
            if (jsonStr === '[DONE]') {
              console.log('收到结束标志 [DONE]')
              // 最后确保全文传回
              if (fullResponseText) {
                onChunk(fullResponseText)
                scrollToBottom()
              }
              onComplete()
              return
            }

            // 跳过空的数据行
            if (!jsonStr) continue

            try {
              const data = JSON.parse(jsonStr)
              
              // 保存 chat_id（如果有）
              if (data.chat_id) {
                chatId.value = data.chat_id
              }

              // ✅ 核心修复逻辑：
              // 只在事件类型为 'conversation.message.delta' 且类型为 'answer' 时才累加
              // 这样就会自动跳过那个包含全文的 'conversation.message.completed' 事件
              if (currentEvent === 'conversation.message.delta' && data.type === 'answer' && data.content) {
                // 只有增量才执行加法
                fullResponseText += String(data.content)
                
                // ✅ 每次更新都把【从头到尾】的全文传出去
                onChunk(fullResponseText)
                scrollToBottom()
              } 
              // 可选：如果想确保最后内容最准，可以在 completed 事件时做一次覆盖
              else if (currentEvent === 'conversation.message.completed' && data.type === 'answer' && data.content) {
                // 强制校准为全文，防止漏字
                fullResponseText = String(data.content)
                onChunk(fullResponseText)
                scrollToBottom()
              }
            } catch (e) {
              // 忽略被截断的 JSON
            }
          }
        }

        // 确保最后的内容被更新
        if (fullResponseText) {
          onChunk(fullResponseText)
          scrollToBottom()
        }

        console.log('流式读取完成，最终内容长度:', fullResponseText.length)
        onComplete()
        return

      } catch (streamError: any) {
        console.error('读取流数据时出错:', streamError)
        onError(`流式读取失败: ${streamError.message || '未知错误'}`)
        return
      } finally {
        // 确保释放读取器
        try {
          reader.releaseLock()
        } catch (e) {
          // 忽略释放错误
        }
      }

    } catch (error: any) {
      console.error('调用本地 API 失败:', error)
      
      // 检查是否是网络错误
      const isNetworkError = error.message?.includes('Failed to fetch') || 
                            error.message?.includes('NetworkError') ||
                            error.message?.includes('ERR_CONNECTION_RESET')
      
      if (isNetworkError) {
        onError('网络请求失败，请检查：\n1. 开发服务器是否正在运行\n2. 网络连接是否正常\n3. 查看控制台了解详细错误信息。')
      } else {
        onError(`请求失败: ${error.message || '未知错误'}`)
      }
      return
    }
}

// 发送消息（流式模式）
const sendMessage = async () => {
  const text = inputText.value.trim()
  if (!text || isLoading.value) return

  // 添加用户消息
  messages.value.push({
    role: 'user',
    content: text,
    timestamp: new Date()
  })

  // 清空输入框
  inputText.value = ''
  
  // 自动调整输入框高度
  if (inputRef.value) {
    inputRef.value.style.height = 'auto'
  }

  // 滚动到底部
  scrollToBottom()

  // 显示加载状态
  isLoading.value = true

  // 创建 AI 消息占位符（用于实时更新）
  const aiMessageIndex = messages.value.length
  messages.value.push({
    role: 'ai',
    content: '',
    timestamp: new Date()
  })

  // 调用流式 API
  await callCozeAPIStream(
    text,
    // onChunk: 接收到数据块时的回调
    // ✅ 修复：直接全量覆盖 AI 消息内容
    // 因为 fullContent 已经是累加好的全文，这里千万不能用 +=
    (fullContent: string) => {
      if (messages.value[aiMessageIndex]) {
        messages.value[aiMessageIndex].content = fullContent
      }
      scrollToBottom()
    },
    // onComplete: 流式传输完成时的回调
    () => {
      isLoading.value = false
      // 确保时间戳已设置
      if (messages.value[aiMessageIndex]) {
        messages.value[aiMessageIndex].timestamp = new Date()
      }
      scrollToBottom()
      console.log('流式传输完成')
    },
    // onError: 发生错误时的回调
    (error: string) => {
      isLoading.value = false
      // 更新错误消息
      if (messages.value[aiMessageIndex]) {
        messages.value[aiMessageIndex].content = error
        messages.value[aiMessageIndex].timestamp = new Date()
      }
      scrollToBottom()
      console.error('流式传输错误:', error)
    }
  )
}

// 监听悬浮窗打开状态，打开时聚焦输入框
watch(isWindowOpen, async (newVal) => {
  if (newVal) {
    await nextTick()
    setTimeout(() => {
      if (inputRef.value) {
        inputRef.value.focus()
      }
      scrollToBottom()
    }, 100)
  }
})

const toggleWindow = () => {
  isWindowOpen.value = !isWindowOpen.value
}

const closeWindow = () => {
  isWindowOpen.value = false
}

// 组件挂载时初始化
onMounted(() => {
  // 可以在这里做一些初始化工作
})
</script>

<style scoped>
.assistant-container {
  position: fixed;
  bottom: 24px;
  right: 24px;
  z-index: 9999;
  pointer-events: none;
}

.floating-ball {
  position: relative;
  width: 80px;
  height: 80px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  box-shadow: 0 8px 24px rgba(102, 126, 234, 0.4);
  transition: all 0.3s ease;
  pointer-events: auto;
  z-index: 10000;
}

.floating-ball:hover {
  transform: scale(1.1);
  box-shadow: 0 12px 32px rgba(102, 126, 234, 0.5);
}

.floating-ball.ball-active {
  transform: scale(1.05);
  box-shadow: 0 8px 24px rgba(102, 126, 234, 0.6);
}

.ball-text {
  color: white;
  font-size: 14px;
  font-weight: 600;
  text-align: center;
  line-height: 1.2;
  user-select: none;
}

.floating-window {
  position: absolute;
  bottom: 100px;
  right: 0;
  width: 400px;
  height: 500px;
  background: white;
  border-radius: 16px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.15);
  pointer-events: auto;
  z-index: 10001;
  overflow: hidden;
}

.window-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 20px;
  border-bottom: 1px solid #e5e7eb;
  background: white;
  border-radius: 16px 16px 0 0;
}

.window-title {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
}

.close-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border: none;
  background: transparent;
  border-radius: 8px;
  cursor: pointer;
  color: #6b7280;
  transition: all 0.2s ease;
}

.close-btn:hover {
  background: #f3f4f6;
  color: #1f2937;
}

.window-content {
  width: 100%;
  height: calc(100% - 65px);
  padding: 0;
  box-sizing: border-box;
  position: relative;
  display: flex;
  flex-direction: column;
  background: #f9fafb;
}

/* 消息列表区域 */
.messages-container {
  flex: 1;
  overflow-y: auto;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.messages-container::-webkit-scrollbar {
  width: 6px;
}

.messages-container::-webkit-scrollbar-track {
  background: transparent;
}

.messages-container::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 3px;
}

.messages-container::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}

/* 欢迎消息 */
.welcome-message {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
  text-align: center;
}

.welcome-icon {
  font-size: 48px;
  margin-bottom: 12px;
}

.welcome-text {
  color: #6b7280;
  font-size: 14px;
  margin: 0;
}

/* 消息项 */
.message-item {
  display: flex;
  gap: 8px;
  align-items: flex-start;
  animation: messageFadeIn 0.3s ease;
}

@keyframes messageFadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.user-message {
  flex-direction: row-reverse;
}

.ai-message {
  flex-direction: row;
}

.message-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  flex-shrink: 0;
  background: #f3f4f6;
}

.user-message .message-avatar {
  background: #667eea;
}

.message-content {
  max-width: 75%;
  display: flex;
  flex-direction: column;
}

.user-message .message-content {
  align-items: flex-end;
}

.ai-message .message-content {
  align-items: flex-start;
}

.message-bubble {
  padding: 10px 14px;
  border-radius: 12px;
  background: white;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
  word-wrap: break-word;
  word-break: break-word;
}

.user-message .message-bubble {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border-bottom-right-radius: 4px;
}

.ai-message .message-bubble {
  background: white;
  color: #1f2937;
  border-bottom-left-radius: 4px;
}

.message-text {
  margin: 0;
  font-size: 14px;
  line-height: 1.5;
  white-space: pre-wrap;
}

.message-time {
  display: block;
  font-size: 11px;
  margin-top: 4px;
  opacity: 0.7;
}

/* 输入区域 */
.input-container {
  padding: 12px 16px;
  background: white;
  border-top: 1px solid #e5e7eb;
}

.input-wrapper {
  display: flex;
  gap: 8px;
  align-items: flex-end;
}

.message-input {
  flex: 1;
  padding: 10px 14px;
  border: 1px solid #e5e7eb;
  border-radius: 20px;
  font-size: 14px;
  font-family: inherit;
  resize: none;
  max-height: 120px;
  overflow-y: auto;
  background: #f9fafb;
  transition: all 0.2s ease;
}

.message-input:focus {
  outline: none;
  border-color: #667eea;
  background: white;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.message-input::placeholder {
  color: #9ca3af;
}

.send-button {
  width: 40px;
  height: 40px;
  border: none;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  transition: all 0.2s ease;
  box-shadow: 0 2px 4px rgba(102, 126, 234, 0.3);
}

.send-button:hover:not(.disabled) {
  transform: scale(1.05);
  box-shadow: 0 4px 8px rgba(102, 126, 234, 0.4);
}

.send-button.disabled {
  opacity: 0.5;
  cursor: not-allowed;
  transform: none;
}

.send-spinner {
  width: 16px;
  height: 16px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.6s linear infinite;
}

/* 打字指示器 */
.typing-indicator {
  display: flex;
  gap: 4px;
  padding: 8px 0;
}

.typing-indicator span {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #9ca3af;
  animation: typing 1.4s infinite;
}

.typing-indicator span:nth-child(2) {
  animation-delay: 0.2s;
}

.typing-indicator span:nth-child(3) {
  animation-delay: 0.4s;
}

@keyframes typing {
  0%, 60%, 100% {
    transform: translateY(0);
    opacity: 0.7;
  }
  30% {
    transform: translateY(-10px);
    opacity: 1;
  }
}

.window-overlay {
  position: fixed;
  inset: 0;
  background: transparent;
  z-index: 10000;
  pointer-events: auto;
}

/* 过渡动画 */
.window-fade-enter-active,
.window-fade-leave-active {
  transition: all 0.3s ease;
}

.window-fade-enter-from {
  opacity: 0;
  transform: translateY(20px) scale(0.95);
}

.window-fade-leave-to {
  opacity: 0;
  transform: translateY(20px) scale(0.95);
}

/* 响应式设计 */
@media (max-width: 640px) {
  .assistant-container {
    bottom: 16px;
    right: 16px;
  }

  .floating-ball {
    width: 64px;
    height: 64px;
  }

  .ball-text {
    font-size: 12px;
  }

  .floating-window {
    width: calc(100vw - 32px);
    height: 60vh;
    bottom: 90px;
    right: 16px;
    left: 16px;
    width: auto;
  }
}

/* 调试信息样式 */
.debug-info {
  margin-top: 12px;
  padding: 12px;
  background: #f3f4f6;
  border-radius: 8px;
  font-size: 11px;
  color: #6b7280;
  border: 1px solid #e5e7eb;
}

.debug-info summary {
  cursor: pointer;
  font-weight: 600;
  color: #4b5563;
  margin-bottom: 8px;
}

.debug-info pre {
  margin: 0;
  padding: 8px;
  background: white;
  border-radius: 4px;
  overflow-x: auto;
  font-size: 10px;
  line-height: 1.4;
  max-height: 200px;
  overflow-y: auto;
}
</style>




