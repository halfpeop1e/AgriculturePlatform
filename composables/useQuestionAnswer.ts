import { useAxios } from "./useAxios"
import { ElMessage } from 'element-plus'  // 补充导入ElMessage
import type { Question, QuestionDetail, Answer } from "~/types/questionAnswer";

const useAxiosInstance = useAxios()

// 获取问答列表
export async function getQuestionList() {
  try {
    const userStore = useUserStore()
    const response = await useAxiosInstance.get<{data:Question[]}>(`/question/list`,{
      params:{
        userId: userStore.userId,
        role: userStore.role
      }
    })
    if (response.status === 200) {
      console.log('获取问答列表成功', response.data)
      return response.data.data
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
    // 修正1：修复模板字符串语法错误（缺少反引号闭合）
    // 修正2：正确使用泛型参数（QuestionDetail是类型，应放在<>中）
    const response = await useAxiosInstance.get(`/question/${id}`)
    if (response.status === 200) {
      console.log('获取问答详情成功', response.data)
      return response.data.data as QuestionDetail
    }
    throw new Error('获取问答详情失败')
  } catch (err) {
    console.error('获取问答详情失败', err)
    throw err
  }
}

// 提交问题
export async function submitQuestion(data: Omit<Question, 'id' | 'date'  | 'answerCount' | 'isAnswered'>) {
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
    const response = await useAxiosInstance.post(`/question/answer/${questionId}`, {
      questionId,
      content
    })
    if (response.status === 200) {
      ElMessage.success('回答提交成功')
      console.log('回答提交成功', response.data)
      return response.data
    }
    throw new Error('回答提交失败')
  } catch (err) {
    console.error('回答提交失败', err)
    throw err
  }
}
