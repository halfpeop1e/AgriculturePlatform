import  {testprofile} from "~/types/test"
export const nav = [
  {
    label: '产品交易',
    items: [
      { label: '所有商品',routername:'/product/allproduct',icon: 'pi pi-fw pi-shopping-cart' },
      { label: '出售商品',routername:'/product/sellproduct',icon: 'pi pi-fw pi-tags' },
      { label: '我的订单',routername:'/product/myorder/my_buy',icon: 'pi pi-fw pi-list' },
      { label: '我的发布',routername:'/product/myrelease',icon: 'pi pi-fw pi-money' },
    ],
  },
  {
    label: '专家问答',
    items: [
      { label: '问题汇总', routername: '/question', icon: 'pi pi-fw pi-comments' },
      { label: '联系专家', routername: '/expert', icon: 'pi pi-fw pi-user-plus' },
      { label: '问答论坛',routername:'/allpost',icon: 'pi pi-fw pi-comments' },
    ],
  },
  {
    label: '金融产品',
    items:[
      {
        label:'贷款产品',routername:'/money/loan',icon: 'pi pi-fw pi-wallet'
      },
      {
        label:'金融产品',routername:'/money/finance/product',icon: 'pi pi-fw pi-chart-line'
      },
    ]
  },
  {
    label: '农业知识',
    items: [
        { label: '农业知识',routername:'/knowledge',icon: 'pi pi-fw pi-book' },
    ],
  },
]
// 导出为纯数据（不包含函数），在运行时由调用方构建带 command 的模型
export const menuData = [
  {
    label: '个人中心',
    icon: 'pi pi-user',
    to: `/profile/${testprofile.userId}`
  },
  {
    label: '设置',
    icon: 'pi pi-cog',
    to: '/setting'
  },
  {
    separator: true
  },
  {
    label: '退出登录',
    icon: 'pi pi-sign-out',
    action: 'logout'
  }
]
