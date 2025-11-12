/*
  myOrder.d.ts
  说明：订单数据结构定义
  字段：
    - orderId: 订单唯一标识
    - name: 商品名称
    - quantity: 购买数量
    - totalprice: 订单总价
    - status: 订单状态（如：待支付、已完成等，字符串）
    - type: 订单类型（'buy' 买入订单 或 'sell' 卖出订单）
*/
export interface Order {
    orderId: string
    name: string
    quantity: number
    totalprice: number
    status: string
    type: 'buy' | 'sell'
}
