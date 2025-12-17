
# AgriculturePlatform

[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/halfpeop1e/AgriculturePlatform)

AgriculturePlatform 是一个基于 Nuxt 3 + Vue 3 的农业服务平台样例，整合了商品展示、下单、订单管理、金融（贷款/融资）申请与审批、知识/问答等模块，旨在作为农业数字化和金融服务的演示与开发基础。

核心目标：为农产品流通与小微农业业务提供一个可扩展的前端模板，包含用户身份、商品发布与管理、订单流程、以及金融审批工作流示例。

## 技术栈

- 框架：Nuxt 3 + Vue 3 (Script Setup + Composition API)
- 状态管理：Pinia
- UI：PrimeVue、Element Plus、Tailwind CSS（混合使用）
- 打包/运行：bun / Vite
- 后端示例：Go (项目内部分示例服务文件)

## 本地开发

安装依赖（推荐使用 `bun`）：

```bash
# bun
bun install

# npm
npm install

# pnpm
pnpm install

# yarn
yarn install
```

启动开发服务器：

```bash
# 使用 bun
bun run dev

# 或者使用 npm
npm run dev
```

## 常用命令

- 启动开发：`bun run dev` 或 `npm run dev`
- 构建生产：`bun run build` 或 `npm run build`
- 本地预览：`bun run preview` 或 `npm run preview`

（若使用其他包管理器，把 `bun` 换为 `pnpm` 或 `yarn`）

## 项目结构（概要）

- `pages/`：Nuxt 页面路由（包含 `homePage.vue`, `login.vue`, `register.vue`, `profile/[userid].vue`, `finaceJustice.vue` 等）
- `components/`：可复用组件（订单卡片、对话框、图片展示等）
- `composables/`：业务级组合函数（`getProduct.ts`, `useLogin.ts` 等）
- `layouts/`：全局布局（`homePageLayout.vue`, `profilePageLayout.vue`, `financePageLayout.vue`）
- `middleware/`：路由守卫（用户登录、角色跳转与页面布局分配）
- `plugins/`：Nuxt 插件（例如用户恢复、全局组件注册等）
- `public/`：静态资源（logo、二维码等）

更多文件请参考仓库树。

## 重要行为与约定

- 身份与路由：项目通过 Pinia 保存用户信息与角色（`useUserStore`），并在 `middleware/userLogin.global.ts` 中按角色做重定向（例如 `finance` 角色跳转到 `/finaceJustice` 并仅可访问审批和个人信息页）。
- 布局控制：部分页面通过中间件在 `to.meta.layout` 上设置布局（`financePageLayout` / `profilePageLayout` 等），以便不同角色使用不同外观。
- 支付二维码：`components/paydialog.vue` 提供带倒计时的支付二维码弹窗示例，倒计时结束会触发 `expired` 事件。

## 配置与环境变量

- 项目会从后端请求数据，若你有自己的后端服务，请在 `.env` 或 Nuxt 配置中设置对应 API 地址（例如 `API_BASE_URL`）。

## 开发注意事项

- Pinia：确保在应用入口正确注册 Pinia（`app.use(pinia)`），否则在服务端/客户端执行 `useUserStore()` 等会报错（错误提示：getActivePinia() was called but there was no active Pinia）。
- 布局键（LayoutKey）：在代码中使用 Nuxt 布局名称时请使用与 `layouts/` 文件名匹配的键（例如 `profilePageLayout`、`financePageLayout`），以避免类型错误。
- 混合 UI：项目同时使用 PrimeVue 与 Element Plus，保持组件样式一致性时请注意样式冲突。

## 贡献与扩展

- 若要增加 API 集成、审批流程或支付通道，请在 `composables/` 中实现数据访问，并在 `pages/` 下新增或扩展页面。
- 建议编写对应的单元测试与端到端测试（项目当前未内置测试框架）。

## 联系

如需帮助或讨论实现细节，请在仓库 issue 中留言或联系仓库维护者。

---
_自动生成：已根据项目源码补充 README，要我把 README 再精简或加入运行截图/部署示例吗？_
