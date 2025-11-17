/*
  comfirmOrder.d.ts
  说明：确认订单的请求接口类型（当前未使用）
  字段：
    - orderId: 订单 ID
    - productId: 产品 ID
    - quantity: 购买数量
    - totalprice: 总价
    - buyer: 买家 ID
    - saler: 卖家 ID（注：拼写应为 seller）
*/
export interface comfirmOrderRequest {
    orderId:string,
    productId:string,
    quantity:number,
    totalprice:number,
    buyer:string
    saler:string
}
