export interface Expert {
  id: string
  name: string
  avatar: string
  field: string // 专业领域
  introduction: string // 简介
  skills: string[] // 擅长技能
  consultCount: number // 咨询次数
  rating: number // 好评率
  responseTime?: string // 响应时间
  recentCases?: {
    date: string
    question: string
    answer: string
  }[] // 近期案例
}

export interface ExpertDetail extends Expert {
  education?: string // 教育背景
  experience?: string // 工作经历
  certification?: string[] // 资质证书
  availableTime?: string[] // 可咨询时间
  serviceScope?: string[] // 服务范围
  price?: number // 咨询费用
}