/*
  routes.d.ts
  说明：导航分组的数据结构（用于侧边栏或顶部导航）
  字段：
    - label: 分组标签（如"首页"、"产品管理"）
    - items: 分组下的导航项数组，每项包含 label（显示名称）、icon（图标 class）、routername（路由路径）
*/
export interface RouteType{
    label: string;
    items: {label: string; icon: string; routername: string}[];
}
