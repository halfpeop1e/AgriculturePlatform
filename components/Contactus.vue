<template>
  <footer class=" py-10 px-8 w-full">
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-6xl mx-auto">
      <!-- 联系信息 -->
      <div class="footer-card">
        <h2 class="footer-title">联系我们</h2>
        <ul class="space-y-2 text-sm text-gray-600">
          <li>📍 地址：北京市朝阳区示例路 88 号</li>
          <li>📞 电话：+86 010-1234-5678</li>
          <li>✉️ 邮箱：contact@example.com</li>
        </ul>
      </div>

      <!-- 快速链接 -->
      <div class="footer-card">
        <h2 class="footer-title">快速链接</h2>
        <ul class="space-y-2 text-sm text-gray-600">
          <li><a href="/" class="hover:text-blue-500 transition-colors">首页</a></li>
          <li><a href="/about" class="hover:text-blue-500 transition-colors">关于我们</a></li>
          <li><a href="/services" class="hover:text-blue-500 transition-colors">服务项目</a></li>
          <li><a href="/contact" class="hover:text-blue-500 transition-colors">在线留言</a></li>
        </ul>
      </div>

      <!-- 社交媒体 -->
      <div class="footer-card">
        <h2 class="footer-title">关注我们</h2>
        <div class="flex space-x-5 mt-2">
          <a href="#" class="hover:text-green-500" aria-label="WeChat">
            <i class="fa-brands fa-weixin text-2xl"></i>
          </a>
          <a href="#" class="hover:text-red-500" aria-label="Weibo">
            <i class="fa-brands fa-weibo text-2xl"></i>
          </a>
          <a href="https://github.com" target="_blank" class="hover:text-gray-700" aria-label="GitHub">
            <i class="fa-brands fa-github text-2xl"></i>
          </a>
        </div>
      </div>
    </div>

    <!-- 底部版权 -->
    <div class="text-center text-gray-400 text-sm mt-8 border-t border-gray-200 pt-4">
      © 2025 助农平台 · All Rights Reserved
    </div>
  </footer>
</template>

<style scoped>
.footer-card {
  @apply bg-white rounded-xl shadow-sm border border-gray-100 p-6 transition-all duration-300 hover:shadow-md hover:-translate-y-1;
}

.footer-title {
  @apply text-lg font-semibold text-gray-800 mb-4 border-l-4 border-green-500 pl-2;
}
</style>
