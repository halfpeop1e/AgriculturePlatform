// https://nuxt.com/docs/api/configuration/nuxt-config
import Aura from '@primeuix/themes/aura';

export default defineNuxtConfig({
   build: {
    transpile: ['vuetify'],
  },
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  css: ['~/assets/css/main.css',
    'primeicons/primeicons.css',
    '~/assets/css/transitions.css',
    'cropperjs/dist/cropper.css'
  ],
  app: {
    pageTransition: { 
      name: 'fade', 
      mode: 'out-in' 
    }
  },
  modules: [
    '@element-plus/nuxt',
    '@pinia/nuxt',
    'nuxt-icons',
    '@primevue/nuxt-module',
    '@nuxtjs/tailwindcss'
  ],
  primevue: {
    options: {
            theme: {
                preset: Aura
            }
        }
  },
  elementPlus: { 

    },
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  typescript: {
    tsConfig: {
      compilerOptions: {
        composite: true,
        noEmit: true // 避免生成多余文件
      }
    }
  },
  nitro: {
    externals: {
      inline: ['@vue/devtools-kit']
    }
  },
  vite: {
    optimizeDeps: {
      exclude: ['@vue/devtools-kit']
    },
    server: {
      proxy: {
        '/coze-api': {
          target: 'https://api.coze.cn',
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/coze-api/, ''),
          configure: (proxy, _options) => {
            proxy.on('error', (err, _req, _res) => {
              console.log('代理错误:', err)
            })
            proxy.on('proxyReq', (proxyReq, req, _res) => {
              console.log('代理请求:', req.method, req.url)
            })
            proxy.on('proxyRes', (proxyRes, req, _res) => {
              console.log('代理响应:', proxyRes.statusCode, req.url)
            })
          }
        }
      }
    }
  },
  // 开发服务器配置
  devServer: {
    port: 3000
  },
  // Runtime Config - 用于存储服务器端配置（保护敏感信息）
  runtimeConfig: {
    // 私有配置（仅在服务器端可用）
    cozeToken: process.env.COZE_TOKEN || 'pat_npKKR94VWrHRCHbON5rHASqxVVSPk6Gq1qBu7M5Lp5JNONxB7jPc6WSGwcYdqbKu',
    cozeBotId: process.env.COZE_BOT_ID || '7585274475898486826',
    // 公共配置（可在客户端使用，但建议只放非敏感信息）
    public: {
      // 如果后端 API 在远程服务器，配置 API 基础 URL
      // 如果为空或未设置，则使用相对路径（同域）
      apiBaseUrl: process.env.NUXT_PUBLIC_API_BASE_URL || ''
    }
  }
})
