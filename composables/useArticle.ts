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