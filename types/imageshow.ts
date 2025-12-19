/*
  imageshow.ts
  说明：首页轮播图配置数据（静态数组）
  结构：每个元素包含 url（图片路径）与 text（覆盖文本）
  使用场景：传递给 imageShow 组件的 images prop
*/
export const Imageitem = [
{ url: '/showImage/img1.jpeg', text: '数字赋能 · 智慧助农' },
{ url: '/showImage/img2.jpeg', text: '买农品 · 助农心' },
{ url: '/showImage/img3.jpeg', text: '产销互联 · 助农共赢' },
{ url: '/showImage/img4.jpg', text: '乡村振兴 · 万众助力' },
]

export const ImageNavigate= [
  { url: '/showImage/nav1.png', text: '农产品', link: '/productList' },
  { url: '/showImage/nav2.png', text: '助农资讯', link: '/newsList' },
  { url: '/showImage/nav3.png', text: '政策法规', link: '/policyList' },
  { url: '/showImage/nav4.png', text: '技术服务', link: '/techList' },
] 
