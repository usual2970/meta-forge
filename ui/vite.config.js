import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import VueDevTools from 'vite-plugin-vue-devtools'

// https://vitejs.dev/config/
export default defineConfig(({mode})=>{

  return {
    plugins: [
      vue(),
      VueDevTools(),
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      }
    },
    base:'/mf',
    server: {
      proxy: mode === 'development'
        ? {
            '/api': {
              target: 'http://127.0.0.1:8090',
              changeOrigin: true,
            },
          }
        : {},
    },
  }
})
