export interface KnowledgeArticle {
  id: string
  title: string
  excerpt: string
  content: string
  image?: string
  tags?: string[]
  author: string
  date: string
}
export interface KnowledgeArticleContent{
    id: string
    title: string
    content: string
    author:string
    date:string
}