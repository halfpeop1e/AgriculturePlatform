/*
  product.d.ts
  说明：商品相关的数据类型
  接口：
    - productResponse: 商品详情（后端返回）— id, name, image, description, price, stock, saler, salerId
    - postProductRequest: 发布商品请求 — 包含 name, images(数组), category, description, price, stock, saler, salerId
  注意：saler 拼写应为 seller（salerId 同理）
*/
export interface productResponse {
    id:string;
    name:string;
    image:string;
    description:string;
    price:number;
    stock:number;
    saler:string;
    salerId:string;
}
export interface postProductRequest {
    name:string;
    images:string[];
    category:string;
    description:string;
    price:number;
    stock:number;
    saler:string;
    salerId:string;
}
