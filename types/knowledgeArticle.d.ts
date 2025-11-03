export interface KnowledgeArticle {
  id: string
  title: string
  excerpt: string
  content: string
  image?: string
  tags?: string[]
  author?: string
  date?: string
}