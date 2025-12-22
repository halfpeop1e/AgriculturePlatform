export interface Author {
  id: string
  nickname: string
  avatar?: string // 可选头像 URL
}

export interface Reply {
  id: string
  content: string
  author: Author
  likes?: number
  liked?: boolean
  createdAt?: string
}

export interface Post {
  id: string
  title: string
  content: string
  author: Author
  likes?: number
  liked?: boolean
  replies?: Reply[]
  createdAt?: string
}

export {}
