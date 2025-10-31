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
  ],
  modules: [
    '@element-plus/nuxt',
    '@pinia/nuxt',
    'nuxt-icons',
    '@primevue/nuxt-module',
    
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
  }
})