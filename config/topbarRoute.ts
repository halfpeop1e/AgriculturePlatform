import  {testprofile} from "~/types/test"
export const nav = [
  {
    label: '产品交易',
    items: [
      { label: '所有商品',routername:'/product/allproduct' },
      { label: '出售/买入',routername:'/homePage'},
      { label: '我的订单',routername:'/homePage' },
    ],
  },
  {
    label: '专家问答',
    items: [
      { label: '问题汇总',routername:'/homePage' },
      { label: '联系专家' ,routername:'/homePage'},
      { label: '提问',routername:'/homePage' },
    ],
  },
  {
    label: '金融产品',
    items:[
      {
        label:'贷款产品',routername:'/homePage'
      },
      {
        label:'金融产品',routername:'/homePage'
      },
    ]
  },
  {
    label: '农业知识',
    items: [
        { label: '农业知识',routername:'/homePage' },
    ],
  },
 
]
export const menuItems = ref([
  {
    label: '个人中心',
    icon: 'pi pi-user',
    command: () => navigateTo(`/profile/${testprofile.userId}`)
 },
  {
    label: '设置',
    icon: 'pi pi-cog',
    command: () => navigateTo('/settings')
  },
  {
    separator: true
  },
  {
    label: '退出登录',
    icon: 'pi pi-sign-out',
    command: () => console.log('退出登录')
  }
])