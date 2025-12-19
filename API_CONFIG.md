# API 配置说明

## 场景说明

根据你的部署架构，选择对应的配置方式：

---

## 场景 1：Nuxt 应用完整部署（前端 + 后端在同一服务器）

**适用情况：**
- Nuxt 应用（包括 `server/api`）部署在同一个服务器上
- 前端和后端在同一个域名下

**配置方式：**
- **无需额外配置**，使用默认的相对路径 `/api/chat`
- 代码会自动使用同域调用

**示例：**
```
前端：https://yourdomain.com
后端：https://yourdomain.com/api/chat  （自动）
```

---

## 场景 2：前端和后端分离部署

**适用情况：**
- 前端部署在 CDN 或静态托管（如 Vercel、Netlify）
- 后端 API 部署在独立的服务器（如 VPS、云服务器）

**配置方式：**

### 方法 A：使用环境变量（推荐）

在项目根目录创建 `.env` 文件：

```env
# 远程后端 API 服务器地址
NUXT_PUBLIC_API_BASE_URL=https://api.yourdomain.com
```

或者在部署平台的环境变量中设置：
- Vercel: 项目设置 → Environment Variables
- Netlify: Site settings → Environment variables
- 其他平台: 参考对应平台的环境变量配置

### 方法 B：修改 nuxt.config.ts

直接在 `nuxt.config.ts` 中设置：

```typescript
runtimeConfig: {
  public: {
    apiBaseUrl: 'https://api.yourdomain.com'  // 你的后端服务器地址
  }
}
```

**示例：**
```
前端：https://yourdomain.com
后端：https://api.yourdomain.com/api/chat
```

---

## 场景 3：本地开发，远程后端

**适用情况：**
- 本地开发前端
- 后端 API 在远程服务器

**配置方式：**

创建 `.env` 文件：

```env
NUXT_PUBLIC_API_BASE_URL=https://api.yourdomain.com
```

或者在 `nuxt.config.ts` 中临时设置（不推荐提交到 Git）。

---

## 后端服务器要求

如果你的后端是独立的服务器（不是 Nuxt Server API），需要确保：

1. **CORS 配置**：允许前端域名访问
2. **API 端点**：提供 `/api/chat` 端点
3. **流式支持**：支持 Server-Sent Events (SSE)
4. **响应头**：
   - `Content-Type: text/event-stream`
   - `Cache-Control: no-cache`
   - `Connection: keep-alive`

### 后端 API 接口规范

**请求：**
```http
POST /api/chat
Content-Type: application/json

{
  "message": "用户消息",
  "user_id": "用户ID"
}
```

**响应：**
- 流式 SSE 格式
- `data: {...}` 格式的数据块
- 结束标志：`data: [DONE]`

---

## 环境变量说明

| 变量名 | 说明 | 必需 | 默认值 |
|--------|------|------|--------|
| `NUXT_PUBLIC_API_BASE_URL` | 远程后端 API 基础 URL | 否 | 空（使用相对路径） |
| `COZE_TOKEN` | Coze API Token | 是 | 代码中默认值 |
| `COZE_BOT_ID` | Coze Bot ID | 是 | 代码中默认值 |

---

## 验证配置

1. 打开浏览器开发者工具（F12）
2. 切换到 Network 标签
3. 发送一条聊天消息
4. 查看请求：
   - **场景 1**：应该看到请求到 `/api/chat`（相对路径）
   - **场景 2/3**：应该看到请求到 `https://api.yourdomain.com/api/chat`（完整 URL）

---

## 常见问题

### Q: 如何判断使用哪种场景？
A: 
- 如果 `server/api/chat.post.ts` 文件存在且部署在同一服务器 → 场景 1
- 如果前端和后端完全分离 → 场景 2
- 如果本地开发但后端在远程 → 场景 3

### Q: 配置后仍然报 CORS 错误？
A: 
- 检查后端服务器的 CORS 配置
- 确认 `NUXT_PUBLIC_API_BASE_URL` 配置正确
- 查看浏览器控制台的完整错误信息

### Q: 如何测试配置是否正确？
A: 
- 查看浏览器 Network 标签中的请求 URL
- 查看服务器日志（如果有）
- 检查响应状态码和内容



