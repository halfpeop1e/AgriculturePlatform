export interface Order {
    orderId: string
    name: string
    quantity: number
    totalprice: number
    status: string
    type: 'buy' | 'sell'
}
