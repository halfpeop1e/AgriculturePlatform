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
  { url: '/navimg/ask.png', text: '咨询专家团队', link: '/expert' },
  { url: '/navimg/finance.jpeg', text: '浏览金融产品', link: '/money/loan' },
  { url: '/navimg/sell.png', text: '进入产品页面', link: '/product/allproduct' },
] 
