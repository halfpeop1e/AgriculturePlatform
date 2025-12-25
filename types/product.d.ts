/*
  product.d.ts
  说明：商品相关的数据类型
  接口：
  - ApiProductRespond: 后端返回的原始商品结构（id 为 number，image 数组为字符串 URL）
  - productResponse: 前端页面使用的商品详情（id 转为 string，image 统一为 { url } 数组）
  - ProductListPayload: 后端返回的分页列表结构（包含分页元数据）
  - ProductListResult: 前端消费的分页结果（list + 分页信息）
  - postProductRequest: 发布商品请求 — 包含 name, images(数组), category, description, price, stock, saler, salerId
  注意：saler 拼写应为 seller（salerId 同理）
*/
export interface ApiProductRespond {
  id: number;
  name: string;
  image: string[];
  description: string;
  price: number;
  stock: number;
  category: string;
  saler: string;
  salerId: string;
}
export interface productResponse {
    id:string;
    name:string;
    image:{url:string}[];
    description:string;
    price:number;
    stock:number;
    saler:string;
    salerId:string;
  category?:string;
}
export interface ProductListPayload {
  productList: ApiProductRespond[];
  total: number;
  page: number;
  pageSize: number;
  hasMore: boolean;
}
export interface ProductListResult {
  list: productResponse[];
  total: number;
  page: number;
  pageSize: number;
  hasMore: boolean;
}
export interface postProductRequest {
    name:string;
    images:string[];
    category:string;
    description:string;
    price:number;
    stock:number;
    contactEmail:string;
    saler:string;
    salerId:string;
}

export interface orderDate {
  dates: [],
  categories: [],
  products: []
}
