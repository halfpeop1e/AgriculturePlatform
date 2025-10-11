// https://nuxt.com/docs/api/configuration/nuxt-config


export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  css: ['~/assets/css/main.css'],
  modules: [
    '@element-plus/nuxt',
    '@pinia/nuxt',
    'nuxt-icons',
  ],
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
  }
})
