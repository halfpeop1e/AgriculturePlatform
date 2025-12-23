import { useAxios } from "./useAxios"
import { ElMessage } from 'element-plus'  // 补充导入ElMessage
import type { Question, QuestionDetail, Answer } from "~/types/questionAnswer";

const useAxiosInstance = useAxios()

// 获取问答列表
export async function getQuestionList() {
  try {
    const response = await useAxiosInstance.get<any>('/question/list')
    if (response.status === 200) {
      console.log('获取问答列表成功', response.data)
      let data = response.data
      // 处理后端可能的包装结构 { code, data, message } 或 { list, total }
      if (data && !Array.isArray(data)) {
        data = data.data || data.list || []
      }
      return Array.isArray(data) ? data : []
    }
    throw new Error('获取问答列表失败')
  } catch (err) {
    console.error('获取问答列表失败', err)
    throw err
  }
}

// 获取问答详情 - 修正语法错误和类型使用错误
export async function getQuestionDetail(id: string) {
  try {
    const response = await useAxiosInstance.get<any>(`/question/${id}`)
    if (response.status === 200) {
      console.log('获取问答详情成功', response.data)
      let data = response.data
      // 如果返回的是 { code, data, message } 结构，提取 data
      if (data && data.data && data.code !== undefined) {
        data = data.data
      }
      return data as QuestionDetail
    }
    throw new Error('获取问答详情失败')
  } catch (err) {
    console.error('获取问答详情失败', err)
    throw err
  }
}

// 提交问题
export async function submitQuestion(data: Omit<Question, 'id' | 'date' | 'author' | 'answerCount' | 'isAnswered'>) {
  try {
    const response = await useAxiosInstance.post<any>('/question/create', data)
    if (response.status === 200) {
      let resData = response.data
      if (resData && resData.data && resData.code !== undefined) {
        resData = resData.data
      }
      return resData
    }
    throw new Error('问题提交失败')
  } catch (err) {
    console.error('问题提交失败', err)
    throw err
  }
}

// 回答问题
export async function submitAnswer(questionId: string, content: string) {
  try {
    const response = await useAxiosInstance.post<any>('/answer/create', {
      questionId,
      content
    })
    if (response.status === 200) {
      let resData = response.data
      if (resData && resData.data && resData.code !== undefined) {
        resData = resData.data
      }
      return resData
    }
    throw new Error('回答提交失败')
  } catch (err) {
    console.error('回答提交失败', err)
    throw err
  }
}
