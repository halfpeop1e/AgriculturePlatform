/*
  useArticle.ts
  说明：与知识库（knowledge）相关的请求封装
  导出函数：
    - getKnowledgeArticleList(): 获取文章列表，返回 KnowledgeArticle[]（成功）
    - getKnowledgeArticleContentById(id): 获取指定文章的完整内容
  错误处理：遇到错误时会在控制台输出并使用 ElMessage 展示错误提示
  依赖：useAxios 返回的 Axios 实例（包含 token 拦截逻辑）
*/
import { useAxios } from "./useAxios"
import type { KnowledgeArticle } from "~/types/knowledgeArticle";
const useAxiosInstance=useAxios()
export async function getKnowledgeArticleList(){
    try{
        const responese =await useAxiosInstance.get<KnowledgeArticle[]>('/knowledge/list')
        if(responese.status===200){
            console.log('获取知识文章列表成功',responese.data)
            return responese.data
        }
        else{
            throw new Error('获取知识文章列表失败')
        }
    }
    catch(err){
        console.error('获取知识文章列表失败',err)
        ElMessage.error('获取知识文章列表失败')
    }
}
export async function getKnowledgeArticleContentById(id:string){
    try{
        const responese =await useAxiosInstance.get(`/knowledge/article/${id}`)
        if(responese.status===200){
            console.log('获取知识文章成功',responese.data)
            return responese.data
        }
        else{
            throw new Error('获取文章具体内容失败')
        }
    }
    catch(err){
        console.error('获取文章具体内容失败',err)
        ElMessage.error('获取文章具体内容失败')
    }
}
