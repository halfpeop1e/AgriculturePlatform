export function formatCurrency(v: number) {
  try {
    return new Intl.NumberFormat('zh-CN', { style: 'currency', currency: 'CNY' }).format(v)
  } catch (e) {
    return `¥ ${v}`
  }
}

export function statusSeverity(status: string) {
  const s = (status || '').toLowerCase()
  if (s.includes('已完成') || s.includes('完成') || s.includes('success') || s.includes('已付')) return 'success'
  if (s.includes('待支付') || s.includes('未支付') || s.includes('pending')) return 'warning'
  if (s.includes('已取消') || s.includes('取消') || s.includes('cancel')) return 'danger'
  return 'info'
}
export function typeSeverity(type: string) {
  const t = (type || '').toLowerCase()
  if (t.includes('buy')) return true
  if (t.includes('sell')) return false
  return true
}
