/*
  test.ts
  说明：测试用的用户资料数据（reactive 对象）
  用途：在开发或单元测试中作为 mock 数据使用，模拟真实的用户资料结构
*/
export const testprofile = reactive({
    userId: 'd54acc27-06f8-44d7-995a-575295ba0704',
    nickname: 'hyw',
    email: 'fengzhu@example.com',
    avatar: '/logo.png',
    bio: '热爱农业科技与前端开发，正在打造智慧农业平台。',
    tags: ['前端工程师', '农业物联网', 'Nuxt3'],
    location: '中国 · 北京',
    joinedAt: '2023-06-18',
    phone: '+86 138-0000-0000',
    address: '北京市通州区智慧农业示范园',
    lastActive: '2025-10-15 14:05'
})
