/*
  knowledgeArticle.d.ts
  说明：知识库文章相关类型定义
  接口：
    - KnowledgeArticle: 完整文章对象（列表+详情），包含 id, title, excerpt, content, tags(可选), author, date
    - KnowledgeArticleContent: 精简版，只包含文章核心内容字段（用于详情页或内容显示）
*/
export interface KnowledgeArticle {
  id: string
  title: string
  excerpt: string
  content: string
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
export interface PostArticleRequest{
    title:string
    excerpt:string
    content:string
    tags?:string[]
    author: string
    date: string
}
