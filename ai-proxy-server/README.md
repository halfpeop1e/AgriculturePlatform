# AI 代理服务器 (BFF)

简易后端代理服务器，专门处理 Coze AI 聊天流式转发。

## 功能特点

- ✅ 解决跨域问题（CORS）
- ✅ 保护 API Token（不暴露给前端）
- ✅ 流式数据传输（实时逐字显示）
- ✅ 轻量级，易于部署和维护

## 快速开始

### 1. 安装依赖

```bash
cd ai-proxy-server
npm install
```

### 2. 配置环境变量

复制 `.env.example` 为 `.env`：

```bash
cp .env.example .env
```

编辑 `.env` 文件，填入你的配置：

```env
PORT=5000
COZE_TOKEN=pat_your_token_here
COZE_BOT_ID=your_bot_id_here
```

### 3. 启动服务器

开发模式（自动重启）：

```bash
npm run dev
```

生产模式：

```bash
npm start
```

### 4. 验证运行

访问健康检查接口：

```
http://localhost:5000/health
```

应该返回：

```json
{
  "status": "ok",
  "message": "AI Proxy Server is running",
  "timestamp": "2025-01-19T..."
}
```

## 前端配置

在你的 Nuxt 项目中，更新 `nuxt.config.ts`：

```typescript
runtimeConfig: {
  public: {
    apiBaseUrl: 'http://localhost:5000'  // 开发环境
    // 生产环境: 'https://api.yourdomain.com'
  }
}
```

或者创建 `.env` 文件：

```env
NUXT_PUBLIC_API_BASE_URL=http://localhost:5000
```

## API 接口

### POST /api/stream-chat

聊天接口，支持流式响应。

**请求：**

```http
POST /api/stream-chat
Content-Type: application/json

{
  "message": "你好",
  "user_id": "user123"
}
```

**响应：**

流式 SSE 格式：

```
data: {"type":"answer","content":"你"}
data: {"type":"answer","content":"好"}
data: [DONE]
```

## 部署

### 本地开发

直接运行 `npm run dev`，服务器会在 `http://localhost:5000` 启动。

### 生产环境

1. **使用 PM2（推荐）**

```bash
npm install -g pm2
pm2 start server.js --name ai-proxy
pm2 save
pm2 startup
```

2. **使用 Docker**

创建 `Dockerfile`：

```dockerfile
FROM node:18-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install --production
COPY . .
EXPOSE 5000
CMD ["node", "server.js"]
```

构建和运行：

```bash
docker build -t ai-proxy-server .
docker run -p 5000:5000 --env-file .env ai-proxy-server
```

3. **使用 systemd（Linux）**

创建服务文件 `/etc/systemd/system/ai-proxy.service`：

```ini
[Unit]
Description=AI Proxy Server
After=network.target

[Service]
Type=simple
User=your-user
WorkingDirectory=/path/to/ai-proxy-server
Environment=NODE_ENV=production
ExecStart=/usr/bin/node server.js
Restart=always

[Install]
WantedBy=multi-user.target
```

启动服务：

```bash
sudo systemctl enable ai-proxy
sudo systemctl start ai-proxy
```

## CORS 配置

默认允许以下来源：

- `http://localhost:3000` (Nuxt 开发服务器)
- `http://localhost:5173` (Vite 开发服务器)

如需添加生产环境域名，编辑 `server.js` 中的 `corsOptions.origin` 数组。

## 故障排查

### 问题 1：端口被占用

修改 `.env` 中的 `PORT` 值，或使用其他端口。

### 问题 2：CORS 错误

检查 `corsOptions.origin` 是否包含你的前端地址。

### 问题 3：流式数据不显示

- 检查 Coze API Token 是否正确
- 查看服务器日志中的错误信息
- 确认网络连接正常

## 日志

服务器会在控制台输出以下日志：

- 收到的请求信息
- Coze API 转发状态
- 流式传输状态
- 错误信息

## 安全建议

1. **生产环境**：使用环境变量存储敏感信息，不要硬编码
2. **HTTPS**：生产环境建议使用 HTTPS
3. **限流**：可以考虑添加请求限流中间件
4. **认证**：可以为接口添加 API Key 认证

## 许可证

MIT

